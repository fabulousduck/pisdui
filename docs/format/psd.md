# PSD file format guide

This document is a combination of a lot of online resources that can be found. This guide aims to take the best from all of them to make a full picture of what a .psd file contains and how it should be parsed and interpreted.

## general structure

The structure of a .psd file is devided into five sections

| sections |
|-------|
| [Header](#header) |
| [Color mode data](#color-mode-data) |
| [Image resources](#image-resources) |
| [Layer and mask information](#layer-and-mask-information) |
| [Image data](#image-data) |

## Notes

 - All lengths are in bytes

 ### Types used for lengths

| length | type |
|:------:|:-----:|
| 1 | `char` aka `byte`|
| 2 | `uint16` aka `short` |
| 3 | `int` |
| 4 | `uin32` aka `long`|
| variable | `[]byte` or `string`|

## Header

    The header of a .psd file is a fixed structure containing
    simple info about the file in general. In the table below
    you can see what exactly is inside of it.

| field | length | description | restricted value|
|:-----:|:------:|:-----------:|:-------------:|
| Signature | 4 |  signature of the file | `"8BPS"` |
| Version | 2 | file version, not to be confused with psd version | `1` |
| Reserved | 6 | buffer reserved for expansion data. | `[0,0,0,0,0,0]` |
| Channels | 2 | amount of channels in the file. | |
| Height | 4 | height of the psd file in pixels | |
| Width | 4 | width of the psd file in pixels |
| Depth | 2 | depth of the psd file |
| Colormode | variable | colormode of the file, see colormodes for possibilities| |

## Color mode data

    The color mode data section here is only really important
    when it is either "Indexed" or "Duotone". all other formats
    have the no special data fields.

    Not much is currently known how the .psd file format uses
    the `duotone` data field and not much description on it can
    be found.



| field | length | description | restricted value |
|:-----:|:------:|:-----------:|:----------------:|
|Length | 4 | length of the following data (0 if not indexed or duotone) | |
|Palette | Length |  palette used by the .psd file | |
|DuotoneData | Length | duotone data used by the .psd file | |

## Image resources

Image resources are blocks of individual data that are global to the .PSD file. These blocks can contain data such as grayscaling, ICC profiles and print info and many more.

A table below shows what data can be found and their tags

## Layer and mask information

## Image data
