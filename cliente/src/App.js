
import React from "react";
import './App.css';
import Customers from './pantallas/Customers'
import Employees from './pantallas/Employees'
import Products from './pantallas/Products'


import {BrowserRouter as Router, Route, Switch, Link} from 'react-router-dom'


export default function router(){
  return(
  <Router>
      <div>
        <ul>
          <li>
            <Link to="/customers">Clientes</Link>
          </li>
          <li>
            <Link to="/employees">Empleados</Link>
          </li>
          <li>
            <Link to="/products">Productos</Link>
          </li>
        </ul>

        <hr />

        {/*
          A <Switch> looks through all its children <Route>
          elements and renders the first one whose path
          matches the current URL. Use a <Switch> any time
          you have multiple routes, but you want only one
          of them to render at a time
        */}
        <Switch>
          <Route exact path="/customers">
            <Customers />
          </Route>
          <Route exact path="/employees">
            <Employees />
          </Route>
          <Route exact path="/products">
            <Products />
          </Route>
          
          
        </Switch>
      </div>
    </Router>
  )
}


