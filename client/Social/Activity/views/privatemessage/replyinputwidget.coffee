class ReplyInputWidget extends ActivityInputWidget

  constructor: (options = {}, data) ->

    options.placeholder    or= ''
    options.inputViewClass or= ReplyInputView

    super options, data
    @setClass 'reply-input-widget'

    @forwardEvent @input, "Enter"


  lockSubmit: -> @locked = yes


  unlockSubmit: -> @locked = no


  empty: -> @input.empty()


  create: ({body, requestData}, callback) ->

    {channel: {id: channelId}}  =  @getOptions()

    {appManager} = KD.singletons
    appManager.tell 'Activity', 'sendPrivateMessage', {channelId, body, requestData}, (err, reply) =>
      return KD.showError err  if err

      callback err, reply.first.lastMessage


  viewAppended: ->

    @addSubView @icon
    @addSubView @input

