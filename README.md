# gosudoc

A Golang client for the Web Services available for the Union Catalog of French University libraries  - SUDOC -  as documented at  [http://documentation.abes.fr/sudoc/manuels/administration/aidewebservices/index.html]

See the (Godoc for this package)[https://godoc.org/github.com/nicomo/gosudoc]

## Covered

- merged
- multiwhere
- frbn2ppn
- ocn2ppn
- dnb2ppn
- Ucatb2ppn
- frcairninfo2ppn
- springerln2ppn
- isbn2ppn
- ean2ppn

## Partially covered

- issn2ppn : we manage all the returned PPNs, but not the additional NoHoldings information associated with some PPNs

## Not covered

- metaperiodical which isn't really an API as it returns a file rather than data
