import React from 'react';
import {HashRouter as Router, Route, Link} from 'react-router-dom';
import AppBar from './AppBar';
import Navigation from './Navigation';
import ReceivedTransactions from './ReceivedTransactions';
import ReceivedTransactionDetails from './ReceivedTransactionDetails';
import SentTransactions from './SentTransactions';

export default class App extends React.Component {
  render() {
    return (
      <Router>
        <div>
          <AppBar />
          <main>
            <Route component={Navigation} />

            <Route exact path="/(sent)?" component={SentTransactions} />
            <Route exact path="/received" component={ReceivedTransactions} />
            <Route path="/received/:id" component={ReceivedTransactionDetails} />
          </main>
        </div>
      </Router>
    );
  }
}
