---
nav_exclude: true
---

# Documentation Structure
here is the template of doc structure:
```bash
+-- ..
|-- other codes
|
|-- docs
|   |-- instructions
|   |   |-- index.md  (parent page)
|   |   |-- demo.md
|   |   |-- integration.md
|   |   |-- migration.md
|   |
|   |-- modules
|   |   |-- index.md      (parent page)
|   |   |-- staking.md
|   |   |-- sequencer.md
|   |
|   |-- (other md files, pages with no children)
|   |-- README.md <- this file
|
|-- ...
```

All markdown files in the repo are rendered as pages in https://docs.decentrio.ventures/rollkit-sdk

# Page syntax

All the Markdown pages in the repo will be included in the doc page alphabetically. However, if we want some pages to be customized, like order, hierarchy,.. we should add some configutations in the begining of the file. Details of all parameters are described below.


```bash
---
nav_order: 1 # Page orders, not required. If not set, the pages will be sorted alphabetically by title.

has_children: true # Not required, default false. If set to true, the page will be parent and have sub-pages: /parent/children

parent: Instructions # Required for children page. If a page is child of a parent page, the param must be set to the title of the parent

nav_exclude: true # Not required, default false. If you dont want to include a MD file in the page, set this param to true
---

# content
...
```