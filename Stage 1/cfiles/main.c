#include <windows.h>
#include <stdio.h>
#include "main.h"

// https://docs.microsoft.com/en-us/windows/desktop/dlls/dynamic-link-library-entry-point-function

BOOL WINAPI DllMain(
    HINSTANCE hinstDLL,  // handle to DLL module
    DWORD fdwReason,     // reason for calling function
    LPVOID lpReserved )  // reserved
{
    // Perform actions based on the reason for calling.
    switch( fdwReason )
    {
        case DLL_PROCESS_ATTACH:
            // Initialize once for each new process.
            // Return FALSE to fail DLL load.
            // printf("[+] Hello from DllMain-PROCESS_ATTACH in Merlin\n");
            // MessageBoxA( NULL, "Hello from DllMain-PROCESS_ATTACH in Merlin!", "Reflective Dll Injection", MB_OK );
            break;

        case DLL_THREAD_ATTACH:
            // Do thread-specific initialization.
            // MessageBoxA( NULL, "Hello from DllMain-PROCESS_ATTACH in Merlin!", "Reflective Dll Injection", MB_OK );
            break;

        case DLL_THREAD_DETACH:
            // Do thread-specific cleanup.
            break;

        case DLL_PROCESS_DETACH:
            // Perform any necessary cleanup.
            break;
    }
    return TRUE;  // Successful DLL_PROCESS_ATTACH.
}

int Magic(){
    Exporter();
    return 0;
}