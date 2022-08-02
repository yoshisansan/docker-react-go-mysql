module local.pkg/sample

replace local.pkg/api => ../api

go 1.18

require local.pkg/api v0.0.0-00010101000000-000000000000
