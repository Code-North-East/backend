## Development Guidelines

---

### Commit standards:
- Commits are meant to be atomic, every single commit will have a small yet significant change.
- The prefix used in a commit message should be aligned with the type of change being made. These are the mentioned types:
    - **feat()** -  for any new feature or enhancement to the codebase.
    - **fix()** - bug fix, in this repo correction related to code examples and snippets.
    - **ci()** - deployment changes wrt GitHub actions or if any Jenkins server being maintained.
    - **chore()** - for any changes which does not change the application behaviour or user-facing changes.
- Commit message should follow this structure, to link each commits with an issue.

    ```
    # type: short message describing the change (refs #ISSUE_ID)
    feat(): Add development guidelines (refs #14)
    ```

### Branch naming conventions:
Branch must be derived from an issue and not be created manually.
