# Skype Enable Attachment

Go Executable to automate enabling skype attachment
On Most of the work machines skype has disabled attachments by default

## Usage

Run this exe to enable skype attachments

## Development

You can enable the attachments by updating following Windows registry for skype

```
HKEY_LOCAL_MACHINE\SOFTWARE\Policies\Skype\Phone
DisableFileTransfer REG_DWORD 0
```

Use build.bat to build binary exe file

```
D:\workspace\skype-enable-attachments>build.bat
```

