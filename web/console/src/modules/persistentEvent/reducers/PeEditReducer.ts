import { combineReducers } from 'redux';
import { reduceToPayload, RecordSet } from '@tencent/qcloud-lib';
import { generateQueryReducer } from '@tencent/qcloud-redux-query';
import { generateFetcherReducer } from '@tencent/qcloud-redux-fetcher';
import * as ActionType from '../constants/ActionType';
import { initValidator } from '../../common/models/Validation';

const TempReducer = combineReducers({
  isOpen: reduceToPayload(ActionType.IsOpenPE, true),

  esAddress: reduceToPayload(ActionType.EsAddress, ''),

  v_esAddress: reduceToPayload(ActionType.V_EsAddress, initValidator),

  indexName: reduceToPayload(ActionType.IndexName, ''),

  v_indexName: reduceToPayload(ActionType.V_IndexName, initValidator)
});

export const PeEditReducer = (state, action) => {
  let newState = state;
  // 销毁设置事件持久化的界面
  if (action.type === ActionType.ClearPeEdit) {
    newState = undefined;
  }
  return TempReducer(newState, action);
};
