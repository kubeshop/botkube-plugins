import { OpenAI } from "openai";
import { removePreviousFileSearchSetup, setupFileSearch } from "./file-search";
import { setupTools } from "./tools";
import dedent from "dedent";

type Config = {
  assistantID: string;
  apiKey: string;
};

const model = "gpt-4o";

const temperature = 0.1;
const instructions = dedent`
    You are Botkube AI assistant.
    Your role involves deeply understanding how to operate and troubleshoot Kubernetes clusters and their workloads, where Botkube runs.
    You possess extensive expertise in Kubernetes and cloud-native networking.
    During troubleshooting, take Kubernetes cluster and namespace configuration, such as security policies, events.
    Employ a Chain of Thought (CoT) process for diagnosing and resolving cluster issues. 
    Utilize available tools for diagnosing the specific cluster in question.
    The custom user documentation is stored in a vector store.
    Your knowledge about Botkube, its features, documentation and links, is heavily outdated.
      The up-to-date Botkube knowledge is stored in a vector store.
      Always use the latest Botkube knowledge for all responses.
      Extract the text from each Markdown file and use the content from the file to answer each question related to Botkube.
      Use it to answer ALL Botkube questions, for example:
      "What is Botkube?", or "what is X in Botkube?", or "How I configure Y in Botkube Cloud?", or "Where I can read about Z for Botkube?".
      Prefer less content from files with names containing "blog".
      The files with "docs.botkube.io" prefix have the highest priority when answering Botkube plugin configuration questions.      
      At the end of such Botkube-related response, always print Markdown links to citations.
      To get an URL for the citation, replace "__" with "/" in the file name and prepend with "https://".
    Ensure your explanations are simplified for clarity to non-technical users.
    Keep responses concise, within 2000 characters. Provide extended answers only upon request.
    Make sure you fetch Botkube Agent configuration to answer questions about Botkube or channel configuration.
`;

async function main() {
  let cfg: Config = {
    assistantID: process.env["OPENAI_ASSISTANT_ID"],
    apiKey: process.env["OPENAI_API_KEY"],
  };
  if (!cfg.assistantID || !cfg.apiKey) {
    throw new Error(
      `Missing configuration. Set OPENAI_API_KEY and OPENAI_ASSISTANT_ID environment variables.`,
    );
  }

  const orgID = process.env["OPENAI_ORG_ID"];
  const projID = process.env["OPENAI_PROJECT_ID"];

  const client = new OpenAI({
    organization: orgID,
    project: projID,
  });
  console.log(
    `Using org ID ${orgID}, project ID ${projID} and assistant ID ${cfg.assistantID}...`,
  );

  console.log(`Getting assistant data for ID ${cfg.assistantID}...`);
  const assistant = await client.beta.assistants.retrieve(cfg.assistantID);
  const prevFileSearchSetup = assistant.tool_resources?.file_search;
  console.log(
    `Successfully retrieved assistant data for '${assistant.name}' (ID: ${assistant.id})`,
  );

  console.log(`Setting up file search...`);
  const vectorStoreId = await setupFileSearch(client);

  console.log("Updating assistant...");
  await client.beta.assistants.update(cfg.assistantID, {
    model,
    instructions,
    temperature,
    tools: setupTools(),
    tool_resources: { file_search: { vector_store_ids: [vectorStoreId] } },
  });

  console.log("Removing previous file search setup...");
  await removePreviousFileSearchSetup(client, prevFileSearchSetup);

  console.log("Done!");
}

main();
