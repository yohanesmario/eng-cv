# Engineering CV

This is a CV template specifically designed for software engineers.

The generated files is in [gen](gen) directory.

## Setup

To set up your environment and release your CV, follow these instructions:

1. Install `wkhtmltopdf`.
    ```bash
    sudo apt install wkhtmltopdf
    ```
2. Install Go by visiting the official Go website and downloading the installer for your operating system. Follow the installation instructions provided.
3. Check [src/config.yaml](src/config.yaml) and edit the values as needed. There's a config for ATS checker too in that file, ***make sure to tailor your keywords uniquely per job application***, as a single keyword mismatch might cause you to not pass the screening process.\
   Examples of possible keyword mismatch:
    1. `Postgres` vs `PostgreSQL`
    2. `MongoDB` vs `Mongo DB`
    3. `Senior Backend Engineer` vs `Senior Back-end Engineer`

## Release

To release your CV, open a terminal and navigate to the root project directory.

***!! Do not commit your changes as the script will handle that for you !!***

Run `./release.sh` to initiate the release process. This script will handle the necessary steps to create a release build.

Alternatively, if you want to create a squash release, run `./squash-release.sh` instead. This script will generate a single, compressed release commit.

After running the release script, the generated files will be placed in the `gen` folder.

That's it! You have successfully set up your environment and can now release your CV using the provided scripts.

---

## YAML Usage

The data source for this cv is in [src/cv.yaml](src/cv.yaml).

### `cv.yaml` Structure

```yaml
full-name: string

contact-info:
  - label: string
    icon: string
    value: string
    uri: URI

summary: string

experiences:
  - title: string
    company: string
    location: string
    start: string
    end: string (optional)
    current: boolean (optional)
    description:
      - string
      - string
    tech-stack:
      - string
      - string

educations:
  - degree: string
    institution: string
    location: string
    start: string
    end: string
```

### Usage Guide

#### Editing YAML

- **Personal Information (`full-name`)**:
   - Modify `full-name` to include the full name of the individual.

- **Contact Information (`contact-info`)**:
   - Edit each entry under `contact-info`:
     - **label**: Specify the type of contact information (e.g., Email, LinkedIn).
     - **icon**: Bootstrap icon (e.g., envelope, linkedin, github, telephone)
     - **value**: Provide the actual contact information (e.g., email address, username).
     - **uri** (optional): If applicable, provide the URI link associated with the contact information.

- **Professional Summary (`summary`)**:
   - Update `summary` to reflect the individual's professional summary. Use `>` for multiline content.

#### Adding Work Experiences

- **Work Experiences (`experiences`)**:
   - To add a new work experience, copy the structure starting from `- title` to `tech-stack`.
   - Fill in details under each field:
     - **title**: Job title.
     - **company**: Company name.
     - **location**: Location of the company.
     - **start**: Start date of employment.
     - **end** (optional): End date of employment.
     - **current** (optional): Set to `true` if currently employed.
     - **description**: List responsibilities and achievements using `-` followed by text.
     - **tech-stack**: List technologies used during the role.

#### Adding Education

- **Education (`educations`)**:
   - To add an educational background, copy the structure starting from `- degree` to `end`.
   - Complete each field:
     - **degree**: Type of degree obtained.
     - **institution**: Name of the educational institution.
     - **location**: Location of the institution.
     - **start**: Year of enrollment or start of education.
     - **end**: Year of graduation or end of education.

## Known Issues

- **Blank PDF**
    - This is caused by a race condition within the `wkhtmltopdf` build pipeline, as of now there's no definitive solution for this, so if you encounter this, run this:
      ```bash
      ./opt/force-build.sh
      # and then
      ./squash-release.sh
      # or
      ./release.sh
      ```
