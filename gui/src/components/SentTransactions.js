import React from 'react';
import Panel from 'muicss/lib/react/panel';
import axios from 'axios';
import {Link} from 'react-router-dom';
import moment from 'moment';
import querystring from 'querystring';
import defaults from 'lodash/defaults';

export default class SentTransactions extends React.Component {
  constructor(props) {
    super(props);
    let query = this.parseQuery(props.location);
    this.state = {loading: true, query, payments: []};
    this.loadPage(query.page);
  }

  componentWillReceiveProps(nextProps) {
    let query = this.parseQuery(nextProps.location);

    if (this.state.query && this.state.query.page == query.page) {
      return;
    }

    this.setState({loading: true, query});
    this.loadPage(query.page);
  }

  parseQuery(location) {
    let search = location.search.substr(1);
    let query = querystring.parse(search);
    return defaults(query, {page: 1});
  }

  loadPage(page) {
    axios.get(`/admin/sent-transactions?page=${page}`)
      .then(response => {
        let loading = false;
        let payments = response.data;
        this.setState({loading, payments});
      })
      .catch(error => this.setState({error: true}));
  }

  render() {
    return <Panel>
        {this.state.error ?
          <div className="mui--text-center">Error loading transactions...</div>
        :
        this.state.loading ?
          <div className="mui--text-center">Loading...</div>
        :
        <div>
          <table className="mui-table">
            <thead>
              <tr>
                <th>ID</th>
                <th>Transaction ID</th>
                <th>Status</th>
                <th>Processed At <i className="material-icons">arrow_drop_down</i></th>
                <th></th>
              </tr>
            </thead>
            <tbody>
            {
              this.state.payments.length == 0 ?
                <tr><td colSpan="5" className="mui--text-center">
                  {this.state.query.page == 1 ? "No transactions found..." : "No more transactions found..."}
                </td></tr>
              :
              this.state.payments.map(tx => {
                let submittedAt = moment(tx.submitted_at);
                return <tr key={tx.id}>
                  <td>{tx.id}</td>
                  <td><a href={"https://horizon.stellar.org/txerations/"+tx.transaction_id} target="_blank">{tx.transaction_id}</a></td>
                  <td>{tx.status}</td>
                  <td>{submittedAt.format()+" ("+submittedAt.fromNow()+")"}</td>
                  <td>
                    <i className="material-icons">zoom_in</i> <a href={"https://www.stellar.org/laboratory/#xdr-viewer?input="+encodeURIComponent(tx.envelope_xdr)+"&type=TransactionEnvelope"} target="_blank">Inspect</a>
                  </td>
                </tr>
              })
            }
            </tbody>
          </table>
          {
            this.state.query.page > 1
            ?
            <div className="mui--pull-left">
              <Link to={{pathname: "/sent", search: `?page=${parseInt(this.state.query.page)-1}`}}>
              <button className="mui-btn mui-btn--flat mui-btn--primary">&laquo; previous page</button>
              </Link>
            </div>
            :
            null
          }
          {
            this.state.payments.length == 10
            ?
            <div className="mui--pull-right">
              <Link to={{pathname: "/sent", search: `?page=${parseInt(this.state.query.page)+1}`}}>
              <button className="mui-btn mui-btn--flat mui-btn--primary">next page &raquo;</button>
              </Link>
            </div>
            :
            null
          }
          <div className="mui--clearfix"></div>
        </div>
        }
      </Panel>
  }
}
