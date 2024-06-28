import { createReadStream, readdirSync } from "node:fs";
import OpenAI from "openai";
import { Assistant } from "openai/resources/beta";

const contentPath = "./content";
const errorStatuses = ["failed", "cancelled"];
const vectorStoreName =
  "Full Botkube knowledge base: documentation, website, blog posts and other content";

export async function setupFileSearch(client: OpenAI): Promise<string> {
  console.log(`Reading directory ${contentPath}...`);
  const files = readdirSync(contentPath);
  if (files.length === 0) {
    throw new Error(`No files found in ${contentPath}`);
  }
  console.log(`Found ${files.length} files.`);
  const fileStreams = files.map((fileName) =>
    createReadStream(`${contentPath}/${fileName}`),
  );

  console.log("Creating vector store...");
  const vectorStore = await client.beta.vectorStores.create({
    name: vectorStoreName,
  });
  console.log(
    `Created vector store '${vectorStore.name}' (ID: ${vectorStore.id})`,
  );

  console.log(
    "Uploading files to vector store and waiting for the file batch processing to complete. This might take a few minutes...",
  );
  const batch = await client.beta.vectorStores.fileBatches.uploadAndPoll(
    vectorStore.id,
    { files: fileStreams },
  );
  console.log(
    `Batch upload completed with status: ${batch.status}; total: ${batch.file_counts.total}, in progress: ${batch.file_counts.in_progress}, completed: ${batch.file_counts.completed}, failed: ${batch.file_counts.failed}`,
  );

  if (errorStatuses.includes(batch.status)) {
    throw new Error(`Batch upload failed with status: ${batch.status}`);
  }

  return vectorStore.id;
}

export async function removePreviousFileSearchSetup(
  client: OpenAI,
  fileSearch?: Assistant.ToolResources.FileSearch,
) {
  const vectorStoreIDs: Array<string> = fileSearch?.vector_store_ids ?? [];
  if (vectorStoreIDs.length === 0) {
    console.log("No previous vector store IDs found.");
    return;
  }

  console.log(`Deleting previous vector store(s)...`);
  for (const id of vectorStoreIDs) {
    while (true) {
      console.log(`- Fetching vector store files page...`);
      const files = await client.beta.vectorStores.files.list(id, {
        limit: 100,
      }); // 100 is max

      if (files.data.length === 0) {
        break;
      }

      console.log(`- Deleting page of ${files.data.length} files...`);
      for (const file of files.data) {
        console.log(`-- Deleting file ${file.id}...`);
        try {
          await client.files.del(file.id);
        } catch (err) {
          console.error(err);
        }
      }
    }

    console.log(`- Deleting vector store ${id}...`);
    await client.beta.vectorStores.del(id);
  }
}
