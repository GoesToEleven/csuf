p = Object.create(HTMLElement.prototype)
p.fly = (e) ->
  e?.preventDefault()
  $(@).append " - Fly!"

p.createdCallback = (e) ->
  console?.log "created!", JSON.stringify(arguments)
  el = $(@)
  el.html(el.html().toUpperCase())

p.attachedCallback = (e) ->
  console?.log "attached!", JSON.stringify(arguments)

p.detachedCallback = (e) ->
  console?.log "detached!", JSON.stringify(arguments)

p.attributeChangedCallback = (e) ->
  console?.log "changed!", JSON.stringify(arguments)

document.registerElement 'kal-el',
  prototype: p
# createdCallback - fires when an instance of the element is created.
# attachedCallback - fires when injected into the document
# detachedCallback - fires when removed from the document
# attributeChangedCallback - fires when an attribute is added, removed, or changed.
$ ->
  $('#main').append('<kal-el>Clark Kent<kal-el')
  $('kal-el').attr('foo', 'bar')

  for el in $('kal-el')
    el.fly()

  setTimeout ->
    $('kal-el').remove()
  , 10000
