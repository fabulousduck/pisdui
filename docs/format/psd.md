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

## Image resources

## Layer and mask information

## Image data