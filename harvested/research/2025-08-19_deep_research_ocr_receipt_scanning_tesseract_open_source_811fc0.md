---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T11:42:36.751792'
profile: deep_research
source: https://tesseract-ocr.github.io/tessdoc/
topic: OCR_receipt_scanning_Tesseract_open_source
---

# OCR_receipt_scanning_Tesseract_open_source

[Skip to the content.](https://tesseract-ocr.github.io/tessdoc/#content)
# Tesseract User Manual
## Tesseract documentation
[View on GitHub](https://github.com/tesseract-ocr/tessdoc)
# Tesseract User Manual
This user manual is for Tesseract versions `5.x`. For versions `4.x.x`, `3.05.02` and older, see the [documentation for old versions](https://tesseract-ocr.github.io/tessdoc/OldVersionDocs.html).
  * [Tesseract User Manual](https://tesseract-ocr.github.io/tessdoc/#tesseract-user-manual)
    * [Introduction](https://tesseract-ocr.github.io/tessdoc/#introduction)
    * [Releases and Changelog](https://tesseract-ocr.github.io/tessdoc/#releases-and-changelog)
    * [Tesseract with LSTM](https://tesseract-ocr.github.io/tessdoc/#tesseract-with-lstm)
    * [5.x.x](https://tesseract-ocr.github.io/tessdoc/#5xx)
      * [Source Code](https://tesseract-ocr.github.io/tessdoc/#source-code)
      * [Binaries](https://tesseract-ocr.github.io/tessdoc/#binaries)
      * [Traineddata Files](https://tesseract-ocr.github.io/tessdoc/#traineddata-files)
      * [Compiling and Installation](https://tesseract-ocr.github.io/tessdoc/#compiling-and-installation)
      * [Usage](https://tesseract-ocr.github.io/tessdoc/#usage)
      * [API Examples](https://tesseract-ocr.github.io/tessdoc/#api-examples)
      * [Technical Information](https://tesseract-ocr.github.io/tessdoc/#technical-information)
      * [Training for Tesseract 5](https://tesseract-ocr.github.io/tessdoc/#training-for-tesseract-5)
      * [Testing](https://tesseract-ocr.github.io/tessdoc/#testing)
      * [External Projects](https://tesseract-ocr.github.io/tessdoc/#external-projects)
    * [User Manual for Old Versions](https://tesseract-ocr.github.io/tessdoc/#user-manual-for-old-versions)


## Introduction
Tesseract is an open source [text recognition (OCR)](https://en.wikipedia.org/wiki/Optical_character_recognition) Engine, available under the [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0).
  * Major version 5 is the current stable version and started with release [5.0.0](https://github.com/tesseract-ocr/tesseract/releases/tag/5.0.0) on November 30, 2021.
  * Newer minor versions and bugfix versions are available from [GitHub](https://github.com/tesseract-ocr/tesseract/releases/).
  * Latest source code is available from [main branch on GitHub](https://github.com/tesseract-ocr/tesseract/tree/main). Open issues can be found in [issue tracker](https://github.com/tesseract-ocr/tesseract/issues), and [planning documentation](https://tesseract-ocr.github.io/tessdoc/Planning.html).


Tesseract can be used directly via [command line](https://tesseract-ocr.github.io/tessdoc/Command-Line-Usage.html), or (for programmers) by using an [API](https://github.com/tesseract-ocr/tesseract/blob/main/include/tesseract/baseapi.h) to extract printed text from images. It supports a [wide variety of languages](https://tesseract-ocr.github.io/tessdoc/Data-Files-in-different-versions.html). Tesseract doesn’t have a built-in GUI, but there are several available from the [3rdParty](https://tesseract-ocr.github.io/tessdoc/User-Projects-%E2%80%93-3rdParty.html) page. External tools, wrappers and training projects for Tesseract are listed under [AddOns](https://tesseract-ocr.github.io/tessdoc/AddOns.html).
Tesseract can be used in your own project, under the terms of the [Apache License 2.0.](http://www.apache.org/licenses/LICENSE-2.0) It has a fully featured API, and can be compiled for a variety of targets including Android and the iPhone. See the [3rdParty](https://tesseract-ocr.github.io/tessdoc/User-Projects-%E2%80%93-3rdParty.html) and [AddOns](https://tesseract-ocr.github.io/tessdoc/AddOns.html) pages for samples of what has been done with it.
If you have a question, first read the [documentation](https://tesseract-ocr.github.io/), particularly the **[FAQ](https://tesseract-ocr.github.io/tessdoc/FAQ.html)** to see if your problem is addressed there. If not, search the [Issues List](https://github.com/tesseract-ocr/tesseract/issues), [Tesseract user forum](http://groups.google.com/group/tesseract-ocr), and if you still can’t find what you need, please ask your question in [Tesseract user forum Google group](http://groups.google.com/group/tesseract-ocr).
Tesseract is free software, so if you want to pitch in and help, please do! If you find a bug and fix it yourself, the best thing to do is to attach the patch to your bug report in the [Issues List](https://github.com/tesseract-ocr/tesseract/issues).
## Releases and Changelog
  * [Release Planning](https://tesseract-ocr.github.io/tessdoc/Planning.html)
  * [API/ABI changes review for Tesseract](https://abi-laboratory.pro/?view=timeline&l=tesseract)
  * [Downloads](https://tesseract-ocr.github.io/tessdoc/Downloads.html)
  * [Releases](https://github.com/tesseract-ocr/tesseract/releases)
  * [Release Notes](https://tesseract-ocr.github.io/tessdoc/ReleaseNotes.html)
  * [Changelog](https://github.com/tesseract-ocr/tesseract/blob/main/ChangeLog)


## Tesseract with LSTM
Tesseract **4.0** added a new OCR engine based on LSTM neural networks. It works well on x86/Linux with official Language Model data available for [100+ languages and 35+ scripts](https://tesseract-ocr.github.io/tessdoc/Data-Files-in-different-versions.html). See [4.0x-Changelog](https://tesseract-ocr.github.io/tessdoc/tess4/4.0x-Changelog.html) for more details.
## 5.x.x
### Source Code
Tesseract **5.x.x** source code is available in the `main` branch of the [repository](https://github.com/tesseract-ocr/tesseract). The `main` branch is using `5.0.0` semver versioning because C++ code modernization caused API incompatibility with 4.x release.
### Binaries
Binaries are available from:
  * [Ubuntu - tesseract-ocr-devel PPA](https://launchpad.net/~alex-p/+archive/ubuntu/tesseract-ocr-devel)
  * [Debian - notesalexp.org](https://notesalexp.org/tesseract-ocr/#tesseract_5.x)
  * [Windows - Tesseract at UB Mannheim](https://github.com/UB-Mannheim/tesseract/wiki)


### Traineddata Files
For detailed information about the different types of models, see [Data Files](https://tesseract-ocr.github.io/tessdoc/Data-Files.html).
Model files for version `4.00` are available from [tessdata tagged 4.00](https://github.com/tesseract-ocr/tessdata/releases/tag/4.00). It has models from November 2016. The individual language file links are available from the following link.
  * [tessdata 4.00 November 2016](https://github.com/tesseract-ocr/tessdoc/blob/master/Data-Files.md#data-files-for-version-400-november-29-2016)


Model files for version `4.0.0` and later are available from [tessdata tagged 4.0.0](https://github.com/tesseract-ocr/tessdata/releases/tag/4.0.0). It has legacy models from September 2017 that have been updated with Integer versions of `tessdata_best` LSTM models. This set of traineddata files has support for both the legacy recognizer with `--oem 0` and for LSTM models with `--oem 1`. These models are available from the following Github repo.
  * [tessdata](https://github.com/tesseract-ocr/tessdata)


Two more sets of `official` traineddata, trained at Google, are made available in the following Github repos. These do not have the legacy models and only have LSTM models usable with `--oem 1`.
  * [tessdata_best](https://github.com/tesseract-ocr/tessdata_best)
  * [tessdata_fast](https://github.com/tesseract-ocr/tessdata_fast)


Language model traineddata files same as listed above for version `4.0.0` can be used with Tesseract `5.x.x`. These are available from:
  * [tessdata](https://github.com/tesseract-ocr/tessdata)
  * [tessdata_best](https://github.com/tesseract-ocr/tessdata_best)
  * [tessdata_fast](https://github.com/tesseract-ocr/tessdata_fast)
  * [tessdata_contrib](https://github.com/tesseract-ocr/tessdata_contrib)
  * [Links to Community Contributions](https://tesseract-ocr.github.io/tessdoc/Data-Files-Contributions.html)


### Compiling and Installation
  * [Compiling and GitInstallation - Linux](https://tesseract-ocr.github.io/tessdoc/Compiling-%E2%80%93-GitInstallation.html)
  * [Compiling - Other O/S](https://tesseract-ocr.github.io/tessdoc/Compiling.html)
  * [Installation](https://tesseract-ocr.github.io/tessdoc/Installation.html)
  * [Docker Containers](https://tesseract-ocr.github.io/tessdoc/Docker-Containers.html)


### Usage
  * [Tips to Improve Recognition](https://tesseract-ocr.github.io/tessdoc/ImproveQuality.html)
  * [Command Line Usage](https://tesseract-ocr.github.io/tessdoc/Command-Line-Usage.html)
  * [Input Formats](https://tesseract-ocr.github.io/tessdoc/InputFormats.html)
  * [Viewer Debugging](https://tesseract-ocr.github.io/tessdoc/ViewerDebugging.html)
  * [Common Errors and Resolutions](https://tesseract-ocr.github.io/tessdoc/Common-Errors-and-Resolutions.html)
  * [Frequently Asked Questions](https://tesseract-ocr.github.io/tessdoc/FAQ.html)


### API Examples
  * [API Example](https://tesseract-ocr.github.io/tessdoc/APIExample.html)
  * [API Example - user_patterns](https://tesseract-ocr.github.io/tessdoc/APIExample-user_patterns.html)
  * [User App Example](https://tesseract-ocr.github.io/tessdoc/User-App-Example.html)
  * [C++ Examples](https://tesseract-ocr.github.io/tessdoc/Examples_C++.html)


### Technical Information
  * [Historical Technical Documentation](https://tesseract-ocr.github.io/tessdoc/tess3/Technical-Documentation.html)
  * [API/ABI changes review for Tesseract](https://abi-laboratory.pro/?view=timeline&l=tesseract)
  * [Manual Pages](https://tesseract-ocr.github.io/tessdoc/Documentation.html#manual-pages)
  * [Source Documentation generated by Doxygen](https://tesseract-ocr.github.io/tessdoc/Documentation.html#source-documentation-generated-by-Doxygen)
  * [Neural Nets in Tesseract](https://tesseract-ocr.github.io/tessdoc/tess4/NeuralNetsInTesseract4.00.html)
  * [VGSL Specs](https://tesseract-ocr.github.io/tessdoc/tess4/VGSLSpecs.html)
  * [VGSL Specs info from Tensorflow](https://github.com/mldbai/tensorflow-models/blob/master/street/g3doc/vgslspecs.md)
  * [Network spec for tessdata_fast models](https://tesseract-ocr.github.io/tessdoc/Data-Files-in-tessdata_fast.html)
  * [Network spec for tessdata_best models](https://tesseract-ocr.github.io/tessdoc/Data-Files-in-tessdata_best.html)
  * [DAS 2016 tutorial slides](https://github.com/tesseract-ocr/docs/tree/master/das_tutorial2016) Slides [#2](https://github.com/tesseract-ocr/docs/blob/main/das_tutorial2016/2ArchitectureAndDataStructures.pdf), [#6](https://github.com/tesseract-ocr/docs/blob/main/das_tutorial2016/6ModernizationEfforts.pdf), [#7](https://github.com/tesseract-ocr/docs/blob/main/das_tutorial2016/7Building%20a%20Multi-Lingual%20OCR%20Engine.pdf) have information about LSTM integration in Tesseract 4.0x.
  * [Tesseract OpenCL - Experimental](https://tesseract-ocr.github.io/tessdoc/TesseractOpenCL.html)


### Training for Tesseract 5
Training with `tesstrain.sh` (a.k.a Tesseract 4 training) is unsupported/abandoned. Please use scripts from [tesseract-ocr/tesstrain](https://github.com/tesseract-ocr/tesstrain) for training.
  * [Train Tesseract LSTM with make from Single Line Images and Groundtruth Transcription](https://github.com/tesseract-ocr/tesstrain)
    * [Examples of Training using tesstrain Makefile](https://github.com/tesseract-ocr/tesstrain/wiki)
  * [Training LSTM Tesseract 5](https://tesseract-ocr.github.io/tessdoc/tess5/TrainingTesseract-5.html) - based on [detailed Tesseract 4 tutorial and guide by Ray Smith](https://tesseract-ocr.github.io/tessdoc/tess4/TrainingTesseract-4.00.html)


### Testing
  * [Benchmarks](https://tesseract-ocr.github.io/tessdoc/Benchmarks.html)
  * [TestingTesseract](https://tesseract-ocr.github.io/tessdoc/TestingTesseract.html)
  * [UNLV Testing of Tesseract](https://tesseract-ocr.github.io/tessdoc/UNLV-Testing-of-Tesseract.html)


### External Projects
  * [AddOns](https://tesseract-ocr.github.io/tessdoc/AddOns.html)
  * [User Projects - 3rdParty](https://tesseract-ocr.github.io/tessdoc/User-Projects-%E2%80%93-3rdParty.html)


### User Manual for Old Versions
  * [Tesseract 4 Documentation](https://tesseract-ocr.github.io/tessdoc/OldVersionDocs.html#tesseract-4)
  * [Tesseract 3 Documentation](https://tesseract-ocr.github.io/tessdoc/OldVersionDocs.html#tesseract-3)

[tessdoc](https://github.com/tesseract-ocr/tessdoc) is maintained by [tesseract-ocr](https://github.com/tesseract-ocr). This page was generated by [GitHub Pages](https://pages.github.com).


## Source Information
- URL: https://tesseract-ocr.github.io/tessdoc/
- Harvested: 2025-08-19T11:42:36.751792
- Profile: deep_research
- Agent: architect
