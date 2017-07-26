import React, { Component } from 'react'
import { Route, IndexRoute, Link } from 'react-router'

// Store
import AppStore from './stores/AppStore'

//Main Component
import App from './components/App'

// Pages
import Blog from './components/Pages/Blog'
import Default from './components/Pages/Default'
import Hobbies from './components/Pages/Hobbies'
import NotFound from './components/Pages/NotFound'


export default (
  <Route path="/" data={AppStore.data} component= {App}>
    <IndexRoute component={Blog}/>
    <Route path="about" component={Default}/>
    <Route path="contact" component={Default}/>
    <Route path="work" component={Hobbies}/>
    <Route path="/work/:slug" component={Hobbies}/>
    <Route path="/blog/:slug" component={Blog}/>

    <Route path="*" component={NotFound}/>
  </Route>
)
