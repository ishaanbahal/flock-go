# flock-go
Flock API SDK in GoLang (Unofficial)

>This repository is currently under development.

Things ready:
 - Methods (All Post APIs)

### TODO: 

- Events
- Token verification
- Get API implementations
- Documentation for usage
- Test cases (verbose)


### Simple Example

```
import channels "github.com/ishaanbahal/flock-go/methods/channels"

func main(){
    channelPost := channels.PostImpl{}
    chanInfo,_ := channelPost.GetInfo(<TOKEN>, <CHANNEL_ID>)
    fmt.Println(chanInfo.Name)
}
```