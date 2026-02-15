## Development Guidelines

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
- Branch must be derived from an issue and not be created manually or locally.
- Once merged with a parent branch (development/main) the private branch must be deleted.

### Epic references:
An epic if, needs to defined in simple words _It can be termed as a document which will contain the collection of tasks which needs to be completed within a milestone_

- A feature should always be part of an epic.
- An **Epic doc** will contain all the necessary design docs related to the development and deployment and integration of that feature.
- A [story](abbreviations.md#Story/Issue/Ticket) should never be derived from a draft epic & once an epic is finalized, stories should be created from that epic.


