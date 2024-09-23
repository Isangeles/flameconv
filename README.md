## Introduction
Flame Converter is a simple CLI tool for converting Flame data input in XML format to the JSON format using native Flame API mappings.

Tool reads the standard input for data in XML format and outputs the converted JSON data string.

## Build
To build use go build command:
```
go build
```

## Usage
Converter requires the type flag specified with the data type of the input to convert.

To convert example XML data file:
```
cat npc.xml | flameconv -type characters
```

## Contact
* Isangeles <<ds@isangeles.dev>>

## License
Copyright 2024 Dariusz Sikora <<ds@isangeles.dev>>

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
MA 02110-1301, USA.
