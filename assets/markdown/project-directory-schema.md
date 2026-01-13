## Project Structure

---
### Directory schema for the backend knowledge base repository.

```
backend/
├── assets/             # all the media files supporting the markdown
├── Cloud/              # documentation directory
│   └── aws/
│   └── gcp/
├── golang/             # source code and documentation directory
├── java/               # source code and documentation directory (Intellij compatible)
├── Projects/           # documentation directory (No business logic, just code snippets of Intellectual property)
├── javascript/               # source code and documentation directory

````
Summarize:

-  Each language directory behaves as a standalone project.
-  Use *Sentence casing* when you have a directory for Documentation.
-  Use *Lower casing* when you have a directory containing source code and documentation.