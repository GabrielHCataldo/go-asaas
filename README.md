go-asaas
=================
<img align="right" src="gopher-asaas.png" alt="">

[![Project status](https://img.shields.io/badge/version-v1.0.2-vividgreen.svg)](https://github.com/GabrielHCataldo/go-asaas/releases/tag/v1.0.2)
[![Go Report Card](https://goreportcard.com/badge/github.com/GabrielHCataldo/go-asaas)](https://goreportcard.com/report/github.com/GabrielHCataldo/go-asaas)
[![Coverage Status](https://coveralls.io/repos/GabrielHCataldo/go-asaas/badge.svg?branch=main&service=github)](https://coveralls.io/github/GabrielHCataldo/go-asaas?branch=main)
[![Open Source Helpers](https://www.codetriage.com/gabrielhcataldo/go-asaas/badges/users.svg)](https://www.codetriage.com/gabrielhcataldo/go-asaas)
[![GoDoc](https://godoc.org/github/GabrielHCataldo/go-asaas?status.svg)](https://pkg.go.dev/github.com/GabrielHCataldo/go-asaas/asaas)
![License](https://img.shields.io/dub/l/vibe-d.svg)

[//]: # ([![build workflow]&#40;https://github.com/GabrielHCataldo/go-asaas/actions/workflows/go.yml/badge.svg&#41;]&#40;https://github.com/GabrielHCataldo/go-asaas/actions&#41;)
[//]: # ([![Source graph]&#40;https://sourcegraph.com/github.com/go-asaas/asaas/-/badge.svg&#41;]&#40;https://sourcegraph.com/github.com/go-asaas/asaas?badge&#41;)
[//]: # ([![TODOs]&#40;https://badgen.net/https/api.tickgit.com/badgen/github.com/GabrielHCataldo/go-asaas/asaas&#41;]&#40;https://www.tickgit.com/browse?repo=github.com/GabrielHCataldo/go-asaas&#41;)

Projeto go-asaas contem todas as interações financeiras com o Gateway de Pagamento Asaas
facilitando a implementação de pagamentos no seu projeto Go.

Instalação
------------

Use go get.

	go get github.com/GabrielHCataldo/go-asaas

Em seguida, importe o pacote go-asaas para seu próprio código.

	import "github.com/GabrielHCataldo/go-asaas/asaas"

Usabilidade e Documentação
------------
Consulte a documentação Asaas https://docs.asaas.com/reference/comece-por-aqui
para obter documentos de uso detalhados.

##### Exemplos:

- [Cliente](https://github/GabrielHCataldo/go-asaas/blob/main/_example/customer/main.go)
- [Cobrança](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/charge/main.go)
- [Assinatura](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/subscription/main.go)
- [Transferências](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/transfer/main.go)
- [Antecipação](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/anticipation/main.go)
- [Parcelamento](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/installment/main.go)
- [Notas Fiscais](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/invoice/main.go)
- [Recarga de celular](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/mobile_phone/main.go)
- [Negativação](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/negativity/main.go)
- [Link de pagamento](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/payment_link/main.go)
- [Conta](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/account/main.go)
- [Subcontas](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/subaccount/main.go)
- [Informações fiscais](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/anticipation/main.go)
- [Consulta Serasa](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/credit_bureau/main.go)
- [Notificação](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/notification/main.go)
- [Webhook](https://github.com/GabrielHCataldo/go-asaas/blob/main/_example/webhook/main.go)

Como contribuir
------
Faça um pull request, ou em caso de algum bug encontrado abra
um Issues.

Licença
-------
Distribuído sob licença MIT, consulte o arquivo de licença dentro do código para obter mais detalhes.