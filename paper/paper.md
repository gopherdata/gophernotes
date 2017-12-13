---
title: 'Gophernotes: the Go kernel for Jupyter and nteract'
tags:
  - go
  - golang
  - jupyter
  - nteract
  - data science
authors:
 - name: Daniel L Whitenack
   orcid: 0000-0003-1155-8889
   affiliation: "1" # (Multiple affiliations must be quoted)
affiliations:
 - name: Pachyderm
   index: 1
date: 13 December 2017
bibliography: paper.bib
---

# Summary

Gophernotes [@Gophernotes] is a Go [@Go] kernel for Jupyter notebooks [@Jupyter_Notebooks] and nteract [@Nteract]. It lets you use Go interactively in a browser-based notebook or desktop application. Use gophernotes to create and share documents that contain live Go code, equations, visualizations and explanatory text. These notebooks, with the live Go code, can then be shared with others via email, Dropbox, GitHub and the Jupyter Notebook Viewer. 

## Functionality

To enable the use of Go in notebooks, one needs to build the gophernotes binary and ensure that this is available in the user's local PATH. A JSON configuration file is then placed in the user's local Jupyter configuration directory. This configuration file references the gophernotes binary and names the kernel `Go`. 

Assuming the above steps have been completed, the `Go` kernel should be recognized by Jupyter or nteract when a Jupyter server or the nteract destop application, respectively, is launched. When a new `Go` notebook is generated in one of those systems, Jupyter or nteract will communicate with gophernotes via ZeroMQ [@ZeroMQ]. In particular, these systems will pass cells of code to the gophernotes kernel. Gophernotes will evaluate these and then pass the result(s) back, which is then displayed on the frontend.

## Ongoing Research and Use

Gophernotes is currently being used by the author to teach concurrency, streaming data analysis, and machine learning techniques with O'Reilly media. Gophernotes has also been made available to students and researchers at Purdue University through JupyterHub. 

# Acknowledgements 

This project utilizes a Go interpreter called gomacro [@Gomacro] under the hood to evaluate Go code interactively. The gophernotes logo was designed by the brilliant Marcus Olsson and was inspired by Renee French's original Go Gopher design. The author thanks Spencer Park for all his very helpful input and contributions. 

# References
