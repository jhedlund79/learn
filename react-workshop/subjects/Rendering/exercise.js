////////////////////////////////////////////////////////////////////////////////
// Exercise:
//
// - render DATA.title in an <h1>
// - render a <ul> with each of DATA.items as an <li>
// - now only render an <li> for mexican food (hint: use DATA.items.filter(...))
// - sort the items in alphabetical order by name (hint: use sort-by https://github.com/staygrimm/sort-by#example)
//
// Got extra time?
// - add a select dropdown to make filtering on `type` dynamic
// - add a button to toggle the sort order
// - Hint: you'll need an `updateThePage` function that calls `render`,
//   and then you'll need to call it in the event handlers of the form controls
////////////////////////////////////////////////////////////////////////////////
import React from 'react'
import ReactDOM from 'react-dom'
import sortBy from 'sort-by'

const DATA = {
  title: 'Menu',
  items: [
    { id: 1, name: 'tacos', type: 'mexican' },
    { id: 2, name: 'burrito', type: 'mexican' },
    { id: 3, name: 'tostada', type: 'mexican' },
    { id: 4, name: 'mushy peas', type: 'english' },
    { id: 5, name: 'fish and chips', type: 'english' },
    { id: 6, name: 'black pudding', type: 'english' }
  ]
}
let foodType = 'mexican'
let sortAscending = true

function updateFoodType(event) {
  foodType = event.target.value
  updateThePage()
}

function toggleSortOrder(event) {
  sortAscending = !sortAscending
  updateThePage()
}

function Menu() {
  const items = DATA.items
    .filter(item => item.type === foodType)
    .sort(sortBy(sortAscending ? 'name' : '-name'))
    .map(item => <li key={item.id} >{item.name}</li>)

  return (
    <div>
      <h1>{DATA.title}</h1>
      <p>
        <select onChange={updateFoodType}>
          <option>mexican</option>
          <option>english</option>
        </select>
        <button onClick={toggleSortOrder}>toggle sort order</button>
      </p>
      <ul>
        {items }
      </ul>
    </div>
  )
}

function updateThePage() {
  ReactDOM.render(<Menu/>, document.getElementById('app'), () => {
    require('./tests').run()
  })
}

updateThePage()
