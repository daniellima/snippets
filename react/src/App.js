import logo from './logo.svg';
import './App.css';
import {useState} from 'react'
import ColorText from './ColorText.js'

function App() {
  const [texts, setTexts] = useState([
    {
      id: 1,
      content: 'The first...'
    }
  ])

  const handleClick = function() {
    texts.push({
      id: texts[texts.length-1].id + 1,
      content: 'Hello!'
    })
    setTexts(texts.slice())
  }

  const handleDelete = function(id) {
    setTexts(texts.filter(t => t.id !== id).slice())
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        {texts.map(text => {
          return <ColorText key={text.id} value={text.content} onDelete={() => handleDelete(text.id)}/>
        })}

        <button
            onClick={handleClick}
        >Click me!</button>
      </header>
    </div>
  );
}

export default App;
