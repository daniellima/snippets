import {useState} from 'react'


export default function ColorText({value, onDelete}) {
    let [colorId, setColorId] = useState(0)

    const color = [
        'green',
        'red',
        'cyan'
    ][colorId]

    const handleClick = () => {
        setColorId((colorId + 1) % 3)
    }

    return <p 
        onClick={handleClick} 
        style={{color: color}}>
            {value}

            <button onClick={onDelete}>X</button>
        </p>
}
