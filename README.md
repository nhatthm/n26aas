# `n26api` as a Service

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/n26aas)](https://github.com/nhatthm/n26aas/releases/latest)
[![Build Status](https://github.com/nhatthm/n26aas/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/n26aas/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/n26aas/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/n26aas)
[![Go Report Card](https://goreportcard.com/badge/github.com/nhatthm/httpmock)](https://goreportcard.com/report/github.com/nhatthm/httpmock)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/nhatthm/n26aas)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

`n26aas` provides `n26api` services using service locator pattern.

## Prerequisites

- `Go >= 1.14`

## Install

```bash
go get github.com/nhatthm/n26aas
```

## Usage

**Examples**

```go
package mypackage

import (
	"github.com/google/uuid"
	"github.com/nhatthm/n26aas"
	"github.com/nhatthm/n26api"
)

func buildClient() (*n26api.Client, error) {
	deviceID := uuid.New()
	c := n26aas.NewClient(deviceID,
		n26api.WithUsername("john.doe"),
		n26api.WithPassword("123456"),
	)

	return c, nil
}
```


### Providers

See https://github.com/nhatthm/n26aas/blob/master/provider.go

## Donation

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
