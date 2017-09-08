import React from 'react';
import Tabs from 'muicss/lib/react/tabs';
import Tab from 'muicss/lib/react/tab';

export default class Navigation extends React.Component {
  constructor({history, location}) {
    super();
    this.state = {};
    this.history = history;
    this.state.selected = this.getActive(location.pathname);
  }

  componentWillReceiveProps(nextProps) {
    const locationChanged = nextProps.location !== this.props.location;

    if (locationChanged) {
      let selected = this.getActive(nextProps.location.pathname);
      this.setState({selected});
    }
  }

  onChange(i, id) {
    this.history.push("/"+id);
  }

  getActive(pathname) {
    let state = {};
    if (pathname.indexOf('/sent') == 0) {
      return 0;
    } else if (pathname.indexOf('/received') == 0) {
      return 1;
    } else {
      return 0;
    }
  }

  render() {
    return <Tabs onChange={this.onChange.bind(this)} defaultSelectedIndex={this.state.selected}>
      <Tab value="sent" label="Sent Transactions"></Tab>
      <Tab value="received" label="Received Transactions"></Tab>
    </Tabs>
  }
}
