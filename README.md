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

## License
-------
    Copyright 2018 Aditya Kamble

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
---
