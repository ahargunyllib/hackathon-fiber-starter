# Code Convention

- Filenames and package names should adhere to snake_case convention
- Variable names and function names should follow camelCase convention
- Contracts defined in [```domain```](./domain/) should use PascalCase convention
- Avoid using single-letter variable names
- All structs implementing contracts in the [```internal```](./internal/) directory should be private, thus following camelCase convention for struct names
- When importing packages, prioritize standard libraries first, followed by third-party libraries, and finally local packages.
Example:
```go
import (
    "fmt"
    "log"
    "time"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/ahargunyllib/hackathon-fiber-starter/pkg/log"
)
```
- For functions with more than 3 parameters, parameters should expand vertically rather than horizontally.
Example:
```go
func Send(
    from string,
    to string,
    title string,
    description string,
    file File
) error {

}
```
- When handling application-generated errors, always log them using [```pkg/log```](./pkg/log/log.go)
- Logging info or errors must follow this convention:
```go
// logging with package defined in pkg/log
log.Info(log.LogInfo{
    "data": data
}, "[File Name in All Caps without .go Extension Separated By Space][Method Name] message")

// example
log.Error(log.LogInfo{
    "error": err.Error(),
}, "[USER REPOSITORY][FetchByEmail] failed to fetch by email")
```
- Handle responses in controllers using [```pkg/helpers/http/response```](./pkg/helpers/http/response/response.go)
- Unit tests must cover all functions defined by contracts in [```domain```](./domain/)

# Commit Message Convention

This website follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

Commit message will be checked using husky and commit lint, you can't commit if not using the proper convention below.

## Format

`<type>(optional scope): <description>`
Example: `feat(pre-event): add speakers section`

### 1. Type

Available types are:

- feat → Changes about addition or removal of a feature. Ex: `feat: add table on landing page`, `feat: remove table from landing page`
- fix → Bug fixing, followed by the bug. Ex: `fix: illustration overflows in mobile view`
- docs → Update documentation (README.md)
- style → Updating style, and not changing any logic in the code (reorder imports, fix whitespace, remove comments)
- chore → Installing new dependencies, or bumping deps
- refactor → Changes in code, same output, but different approach
- ci → Update github workflows, husky
- test → Update testing suite, cypress files
- revert → when reverting commits
- perf → Fixing something regarding performance (deriving state, using memo, callback)
- vercel → Blank commit to trigger vercel deployment. Ex: `vercel: trigger deployment`

### 2. Optional Scope

Labels per page Ex: `feat(pre-event): add date label`

\*If there is no scope needed, you don't need to write it

### 3. Description

Description must fully explain what is being done.

Add BREAKING CHANGE in the description if there is a significant change.

**If there are multiple changes, then commit one by one**

- After colon, there are a single space Ex: `feat: add something`
- When using `fix` type, state the issue Ex: `fix: file size limiter not working`
- Use imperative, and present tense: "change" not "changed" or "changes"
- Don't use capitals in front of the sentence
- Don't add full stop (.) at the end of the sentence

# API Naming Convention

For API naming conventions, refer to the following guide: [API Naming Convention](https://restfulapi.net/resource-naming/) and prefix all endpoints with ```/api/v1```

# Pushing Changes

After committing your changes, push them to your personal branch first. Your personal branch name should be your nickname. Subsequently, you can submit a pull request to the dev branch. The owner will merge the dev branch into the master branch once it passes all unit tests and is ready for deployment on VPS.

# References
- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
- [Theodorus Clarence's Conventional Commit Readme](https://theodorusclarence.com/shorts/conventional-commit-readme)
