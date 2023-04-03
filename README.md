# HarvestOvertime

## Build instructions - compile for current OS
1. Install golang (v1.18 or later)
2. Install fyne
    - ```go get fyne.io/fyne/v2@latest```
    - ```go get fyne.io/fyne/v2/cmd/fyne@latest```
3. Compile project
    - ```fyne package -os {darwin (for mac) | linux | windows}```


## Build instructions - cross compile
1. Install golang and fyne
2. Install docker or podman
3. Install fyne-cross
    - ```go install github.com/fyne-io/fyne-cross@latest```
4. Compile project
    - ```fyne-cross darwin -arch=amd64``` For mac x86
    - ```fyne-cross linux -arch=amd64``` For linux x86
    - ```fyne-cross windows -arch=amd64``` For windows x86
5. Grab executable from fyne-cross/bin or fyne-cross/dist


## Resources
- Fyne docs
    - https://developer.fyne.io/index.html
