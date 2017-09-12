import React from 'react';
import Panel from 'muicss/lib/react/panel';
import axios from 'axios';
import moment from 'moment';
import querystring from 'querystring';

export default class ReceivedTransactionDetails extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
    this.id = props.match.params.id;
    this.history = props.history;
    axios.get('/admin/received-payments/'+this.id)
      .then(response => this.setState({response: response.data}))
      .catch(error => this.setState({error: true}));
  }

  asset(asset_type, asset_code, asset_issuer) {
    let asset;
    if (asset_type == 'native') {
      return "XLM";
    } else {
      return <span>
        {asset_code} (<code>{asset_issuer}</code>)
      </span>
    }
  }

  amount(operation) {
    return <span>{operation.amount} {this.asset(operation.asset_type, operation.asset_code, operation.asset_issuer)}</span>;
  }

  back() {
    if (this.history.length == 1) {
      this.history.push("/received");
    } else {
      this.history.goBack();
    }
  }

  reprocess() {
    let force = false;
    if (!window.confirm("Are you sure you want to reprocess this transaction?")) { 
      return;
    }

    if (this.state.response.payment.status == "Success") {
      if (!window.confirm("The payment was successfully processed before, do you want to continue?")) { 
        return;
      }
      force = true;
    }

    this.setState({reprocessing: true});

    axios.post('/reprocess/', querystring.stringify({operation_id: this.state.response.operation.id, force}))
      .then(response => location.reload())
      .catch(error => {
        alert("Error sending reprocess request");
        this.setState({reprocessing: false});
      });
  }

  render() {
    let processedAt;
    if (this.state.response) {
      processedAt = moment(this.state.response.payment.processed_at);
    }

    return <Panel>
      {this.state.error ?
        <div className="mui--text-center">Error loading payment...</div>
      :
      !this.state.response ?
        <div className="mui--text-center">Loading...</div>
      :
        <div>
          <div className="mui--pull-left">
            <button className="mui-btn mui-btn--flat mui-btn--primary" onClick={this.back.bind(this)}>&laquo; back</button>
          </div>
          <div className="mui--pull-right">
            <button className="mui-btn mui-btn--flat mui-btn--danger" onClick={this.reprocess.bind(this)} disabled={this.state.reprocessing}>
            {!this.state.reprocessing ? "Reprocess" : "Reprocessing..."}
          </button>
          </div>
          <div className="mui--clearfix"></div>

          <div className="mui--text-headline">General</div>
          <table className="mui-table details">
            <tbody>
              <tr>
                <td>ID</td>
                <td>{this.state.response.payment.id}</td>
              </tr>
              <tr>
                <td>Status</td>
                <td>{this.state.response.payment.status}</td>
              </tr>
              <tr>
                <td>Processed At</td>
                <td>{processedAt.format()+" ("+processedAt.fromNow()+")"}</td>
              </tr>
            </tbody>
          </table>
          <div className="mui--text-headline">Operation Details</div>
          <table className="mui-table details">
            <tbody>
              <tr>
                <td>Operation ID</td>
                <td><a href={"https://horizon.stellar.org/operations/"+this.state.response.operation.id} target="_blank">{this.state.response.operation.id}</a></td>
              </tr>
              <tr>
                <td>Type</td>
                <td>{this.state.response.operation.type}</td>
              </tr>
              <tr>
                <td>From</td>
                <td><code>{this.state.response.operation.from}</code></td>
              </tr>
              <tr>
                <td>To</td>
                <td><code>{this.state.response.operation.to}</code></td>
              </tr>
              <tr>
                <td>Amount</td>
                <td>{this.amount(this.state.response.operation)}</td>
              </tr>
              <tr>
                <td>Memo</td>
                <td>{this.state.response.operation.memo.memo_type == 'none' ? <i>None</i> : <code>{this.state.response.operation.memo.memo}</code>}</td>
              </tr>
            </tbody>
          </table>
          <div className="mui--text-headline">Compliance Data</div>
          {this.state.response.auth_data ?
            <table className="mui-table details">
              <tbody>
                <tr>
                  <td>Sender</td>
                  <td>{this.state.response.auth_data.sender}</td>
                </tr>
                <tr>
                  <td>Need Info</td>
                  <td>{this.state.response.auth_data.need_info ? "true" : "false"}</td>
                </tr>
                <tr>
                  <td>Tx</td>
                  <td>
                    <code style={{wordBreak: "break-all"}}>{this.state.response.auth_data.tx}</code>
                    <br />
                    <i className="material-icons">zoom_in</i> <a href={"https://www.stellar.org/laboratory/#xdr-viewer?input="+encodeURIComponent(this.state.response.auth_data.tx)+"&type=Transaction"} target="_blank">Inspect</a>
                  </td>
                </tr>
                <tr>
                  <td>Attachment</td>
                  <td><pre>{JSON.stringify(JSON.parse(this.state.response.auth_data.attachment), null, 2)}</pre></td>
                </tr>
              </tbody>
            </table>
            :
            <div>No data...</div>
          }
        </div>
      }
      </Panel>
  }
}
