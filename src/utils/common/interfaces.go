package common

import gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"

type MessageHandler func(gossiper.AMQMessage) any
