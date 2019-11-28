local fio = require('fio')
local msgpack = require('msgpack')
local json = require('json')

local file = fio.open('msgpack', {'O_RDWR', 'O_CREAT'})

local object = {
    { {1, 2, 3, 4, 5} },
    'greeting',
}

file:write(msgpack.encode(object))
file:close()

fio.chmod('msgpack', tonumber('0777', 8))

print(json.encode(object))
