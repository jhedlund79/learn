////////////////////////////////////////////////////////////////////////////////
// Exercise:
//
// - change the contents of the render function and save the file
// - see the updates automatically in your browser without refreshing!
////////////////////////////////////////////////////////////////////////////////
import React from 'react'
import ReactDOM from 'react-dom'


const view = {
  people: [
     { firstName: 'Jason', lastName:'Hedlund' },
     { firstName: 'Joe', lastName: 'Smoe' }
  ]
}

function ContactList () {

  return (
    <div className="hot">
      {view.people.map((person, index) => (
       <div>{index +1} {person.firstName} {person.lastName}</div>
      ))}
    </div>
  )
}

ReactDOM.render(<ContactList/>, document.getElementById('app'))
