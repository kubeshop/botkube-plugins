import { createReadStream, readdirSync } from "node:fs";
import OpenAI from "openai";

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

  console.log("Checking if vector store already exists...");
  const vectorStores = await client.beta.vectorStores.list();
  const vsToDelete = vectorStores.data.find(
    (vs) => vs.name === vectorStoreName,
  );
  if (vsToDelete) {
    console.log(`Found vector store ${vsToDelete.id}. Deleting...`);
    await client.beta.vectorStores.del(vsToDelete.id);
  }

  console.log("Creating vector store...");
  const vectorStore = await client.beta.vectorStores.create({
    name: vectorStoreName,
  });

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
