const express = require('express');
require('dotenv').config();
const app = express();
const cors = require('cors');
app.use(cors());
app.use(express.json())


app.get("/api/serverDaniel", async(req, res) => {

    
  
    return res.status(200).json({message:'Hola mundo'});

});

 
const PORT = process.env.PORT || 3000;
const server = app.listen(PORT, () => console.log(`Listening on port: ${PORT}`));

function closeServer() {
  server.close();
}  

module.exports = { app, closeServer };