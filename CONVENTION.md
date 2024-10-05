### Code Convention

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

### Commit Convention

All commit messages should adhere to [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)

### API Naming Convention

For API naming conventions, refer to the following guide: [API Naming Convention](https://restfulapi.net/resource-naming/) and prefix all endpoints with ```/api/v1```

### Pushing Changes

After committing your changes, push them to your personal branch first. Your personal branch name should be your nickname. Subsequently, you can submit a pull request to the dev branch. The owner will merge the dev branch into the master branch once it passes all unit tests and is ready for deployment on VPS.
