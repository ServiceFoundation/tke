import * as React from 'react';
import { RootProps } from '../ClusterApp';
import { RegionBar } from '../../../common/components';
import { bindActionCreators } from '@tencent/qcloud-lib';
import { allActions } from '../../actions';
import { connect } from 'react-redux';
import { Button, Text, RegionPanel, RegionOption, RegionSelect, Justify } from '@tea/component';
import { FetchState } from '@tencent/qcloud-redux-fetcher';
import { t, Trans } from '@tencent/tea-app/lib/i18n';

const mapDispatchToProps = dispatch =>
  Object.assign({}, bindActionCreators({ actions: allActions }, dispatch), {
    dispatch
  });

@connect(
  state => state,
  mapDispatchToProps
)
export class ClusterHeadPanel extends React.Component<RootProps, any> {
  componentDidMount() {
    const { actions } = this.props;
    // 进行地域的拉取
    actions.region.applyFilter({});
  }

  render() {
    // let { region, actions } = this.props;

    // let content: JSX.Element;

    // if (region.list.fetchState === FetchState.Fetching || region.list.fetched === false) {
    //   content = (
    //     <div style={{ display: 'inline-block', marginTop: '5px' }}>
    //       <i className="n-loading-icon" />
    //       &nbsp; <span className="text">{t('加载中...')}</span>
    //     </div>
    //   );
    // } else {
    //   const { Column, Group } = RegionPanel;
    //   content = (
    //     <RegionSelect
    //       value={region.selection ? region.selection.value + '' : ''}
    //       onChange={value => {
    //         actions.region.selectRegion(+value);
    //       }}
    //     >
    //       <RegionPanel>
    //         <Column>
    //           <Group name="中国大陆">
    //             {region.list.data.records.map((item, index) => {
    //               return (
    //                 <RegionOption key={index} value={item.value + ''} country="China">
    //                   {item.name}
    //                 </RegionOption>
    //               );
    //             })}
    //           </Group>
    //         </Column>
    //       </RegionPanel>
    //     </RegionSelect>
    //   );
    // }

    return (
      <Justify
        left={
          <React.Fragment>
            <h2>{t('集群管理')}</h2>
            {/* {content} */}
          </React.Fragment>
        }
      />
    );
  }
}
