#
# Copyright 2012 James Helferty. All Rights Reserved.
#
# Generates DDS files of various formats for testing purposes.
# Uses June 2010 DirectX SDK on x64 arch

import subprocess
import os


def DDSConvert(texconv, format_list, inputfile, miplevels):
    # Given a space-delimited string of format names, generates dds files for each
    # e.g., texconv -f R8G8B8 -m mipmaps -ft DDS -nologo test.bmp
    for format in formats.split():
        cmd = [texconv, '-f', format, '-m', miplevels, '-ft', 'DDS', '-nologo', inputfile+'.bmp']
        print ' '.join(cmd)
        subprocess.call(cmd)
        os.rename(inputfile+".dds", inputfile+format+".dds")


# DX9
formats_unorm = \
    "R8G8B8 A8R8G8B8 X8R8G8B8 R5G6B5 X1R5G5B5 A1R5G5B5 A4R4G4B4 R3G3B2 " \
    "A8 A8R3G3B2 X4R4G4B4 A2B10G10R10 A8B8G8R8 X8B8G8R8 G16R16 A2R10G10B10 " \
    "A16B16G16R16"

formats_signed = \
    "V8U8 L6V5U5 X8L8V8U8 Q8W8V8U8 V16U16 A2W10V10U10"

formats_compressed = \
    "DXT1 DXT2 DXT3 DXT4 DXT5"

formats_unused = \
    "UYVY R8G8_B8G8 YUY2 G8R8_G8B8 " \
    "D16_LOCKABLE D32F_LOCKABLE " \
    "L16 Q16W16V16U16 R16F G16R16F A16B16G16R16F " \
    "R32F G32R32F A32B32G32R32F " \
    "CxV8U8 " \
    "A8P8 P8 L8 A8L8 A4L4"

# DX10
formats_dx10 = \
    "R32G32B32A32_FLOAT R32G32B32A32_UINT R32G32B32A32_SINT " \
    "R32G32B32_FLOAT R32G32B32_UINT R32G32B32_SINT R16G16B16A16_FLOAT " \
    "R16G16B16A16_UNORM R16G16B16A16_UINT R16G16B16A16_SNORM " \
    "R16G16B16A16_SINT R32G32_FLOAT R32G32_UINT R32G32_SINT " \
    "R10G10B10A2_UNORM R10G10B10A2_UINT R11G11B10_FLOAT R8G8B8A8_UNORM " \
    "R8G8B8A8_UNORM_SRGB R8G8B8A8_UINT R8G8B8A8_SNORM R8G8B8A8_SINT " \
    "R16G16_FLOAT R16G16_UNORM R16G16_UINT R16G16_SNORM R16G16_SINT " \
    "R32_FLOAT R32_UINT R32_SINT R8G8_UNORM R8G8_UINT R8G8_SNORM " \
    "R8G8_SINT R16_FLOAT R16_UNORM R16_UINT R16_SNORM R16_SINT " \
    "R8_UNORM R8_UINT R8_SNORM R8_SINT A8_UNORM R9G9B9E5_SHAREDEXP " \
    "R8G8_B8G8_UNORM G8R8_G8B8_UNORM BC1_UNORM BC1_UNORM_SRGB BC2_UNORM " \
    "BC2_UNORM_SRGB BC3_UNORM BC3_UNORM_SRGB BC4_UNORM BC4_SNORM " \
    "BC5_UNORM BC5_SNORM"

# DX11
formats_dx11 = \
    "B8G8R8A8_UNORM B8G8R8X8_UNORM " \
    "B8G8R8A8_UNORM_SRGB B8G8R8X8_UNORM_SRGB BC6H_UF16 BC6H_SF16 " \
    "BC7_UNORM BC7_UNORM_SRGB"

texconv="C:\Program Files (x86)\Microsoft DirectX SDK (June 2010)\Utilities\\bin\\x64\\texconv"
texconvex="C:\Program Files (x86)\Microsoft DirectX SDK (June 2010)\Utilities\\bin\\x64\\texconvex"
miplevels="4"
inputfile="test"

formats = ' '.join([formats_unorm, formats_compressed])
DDSConvert(texconv, formats, inputfile, miplevels)

#newformats = formats_dx10
#DDSConvert(texconvex, newformats, inputfile, miplevels)

