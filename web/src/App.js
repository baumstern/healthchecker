import React, { useState, useEffect } from 'react';

import './App.css';

import Watch from './client'


function Row({ network, block, timestamp }) {  
  return (
    <div>
    {network}'s last block is {block} (at {timestamp})
    </div>
  );
}

function App() {
  const [ethBlock, setEthBlock] = useState();
  const [klayBlock, setKlayBlock] = useState();

  const [ethTimestamp, setEthTimestamp] = useState();
  const [klayTimestamp, setKlayTimestamp] = useState();

  useEffect(() => {
    var ethIntervalID = setInterval(() => {
      Watch("ethereum").then((response) => {
          setEthBlock(response.block_num);
          setEthTimestamp(response.timestamp);
      })
    }, 5000);

    var klayIntervalID = setInterval(() => {
      Watch("klaytn").then((response) => {
          setKlayBlock(response.block_num);
          setKlayTimestamp(response.timestamp);
      })
    }, 5000);
    return () => {
      clearInterval(ethIntervalID);
      clearInterval(klayIntervalID);
    }; 
  }
  );
  

  return (
    <div className="App">
      <header className="App-header">
        <h1>healthchecker</h1>
        <Row network='ethereum' block={ethBlock} timestamp={ethTimestamp}/>
        <Row network='klaytn' block={klayBlock} timestamp={klayTimestamp}/>
      </header>   
    </div>
  );
}

export default App;
