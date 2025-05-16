module carddeck

go 1.23.1

require (
	local/item v0.0.0
	local/model v0.0.0
)

replace (
	local/item => ../item
	local/model => ../model
)
