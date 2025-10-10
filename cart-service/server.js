const express = require('express')
const app = express()
app.use(express.json())

let cart = []

app.get('/', (req, res) => {
  res.send({ message: 'Welcome to Cart Service' })
})

app.get('/cart', (req, res) => {
  res.json(cart)
})

app.post('/cart', (req, res) => {
  cart.push(req.body)
  res.json({ message: 'Item added', item: req.body })
})

app.delete('/cart/:itemId', (req, res) => {
  const id = parseInt(req.params.itemId)
  cart = cart.filter(i => i.id !== id)
  res.json({ message: 'Item removed' })
})

app.listen(8001, () => {
  console.log('Cart Service listening on 8001')
})