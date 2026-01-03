# Testing Fyne

## Installing on Windows

### Instruction
- Follow the https://docs.fyne.io/started/quick/ for how to install.
- On windows, you need to install C compiler. Instead of using the MSYS2 from the instruction, use MinGW to install directly on Windows. 
    - Check https://dev.to/gamegods3/how-to-install-gcc-in-windows-10-the-easier-way-422j
    - However, don't use the download page from that page, as it has the outdated 32-bit. Instead, use the 64-bit here https://winlibs.com/
        - Unzip and extract `mingw64` folder to `C:\Downloaded Programs`.
        - Add to Windows environment path: `C:\Downloaded Programs\mingw64\bin`.
- This app is created from the documentation on https://docs.fyne.io/started/hello/

### Development
- Edit main.go to desired coding.
- To build, run `go build .`.
- To run on windows, run `go run .`.
- To run on browser, run `fyne serve --icon myapp.png`.

### Testing
- On root folder, run` go test ./...`

### Packaging
- For packaging on Windows, run `fyne package -os windows -icon myapp.png`.
    - To reduce memory footprint, remove debug symbols, run `fyne package -os windows -icon myapp.png -release`.
    - To install to Windows system, run `fyne install -icon myapp.png`.
- For packaging on browser, run `fyne package -os web -icon myapp.png`.
    - A `wasm` folder will be created.
    - Run `cd wasm`.
    - Run `serve`. Localhost of port 3000 should be served.

### References
- For fyne documentation: https://docs.fyne.io/started/quick/
- Installing C compiler directly on Windows: https://winlibs.com/
- [NOT WORKING] For installing C compiler directly on Windows: https://dev.to/gamegods3/how-to-install-gcc-in-windows-10-the-easier-way-422j
- [NOT WORKING] MinGW for 64-bit: https://sourceforge.net/projects/mingw-w64/