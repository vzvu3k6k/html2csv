# html2csv

`html2csv` is a command-line tool that converts HTML tables into CSV format.

Still work in progress.

## Usage

```
$ cat testdata/simple.html
<table>
  <thead>
    <tr>
      <td>name</td>
      <td>price</td>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>apple</td>
      <td>100</td>
    </tr>
    <tr>
      <td>banana</td>
      <td>200</td>
    </tr>
  </tbody>
</table>
$ html2csv < testdata/simple.html
name,price
apple,100
banana,200
```

## Installation

```sh
go install github.com/vzvu3k6k/html2csv@latest
```

## Development

```sh
go test
```
