# Testing Fyne

## Installing on Windows

### Instruction
- Follow the https://docs.fyne.io/started/quick/ for how to install.
- On windows, you need to install C compiler. Instead of using the MSYS2 from the instruction, use MinGW to install directly on Windows. 
    - Check https://dev.to/gamegods3/how-to-install-gcc-in-windows-10-the-easier-way-422j
    - However, don't use the download page from that page, as it has the outdated 32-bit. Instead, use the 64-bit here https://winlibs.com/
        - Unzip and extract `mingw64` folder to `C:\Downloaded Programs`.
        - Add to Windows environment path: `C:\Downloaded Programs\mingw64\bin`.

### References
- For fyne documentation: https://docs.fyne.io/started/quick/
- Installing C compiler directly on Windows: https://winlibs.com/
- [NOT WORKING] For installing C compiler directly on Windows: https://dev.to/gamegods3/how-to-install-gcc-in-windows-10-the-easier-way-422j
- [NOT WORKING] MinGW for 64-bit: https://sourceforge.net/projects/mingw-w64/