// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'message_api_client.dart';

// **************************************************************************
// RetrofitGenerator
// **************************************************************************

class _MessageApiClient implements MessageApiClient {
  _MessageApiClient(this._dio, {this.baseUrl}) {
    ArgumentError.checkNotNull(_dio, '_dio');
    baseUrl ??= 'http://localhost:8080';
  }

  final Dio _dio;

  String baseUrl;

  @override
  Future<Map<String, List<Message>>> getMessages() async {
    const _extra = <String, dynamic>{};
    final queryParameters = <String, dynamic>{};
    final _data = <String, dynamic>{};
    final _result = await _dio.request<Map<String, dynamic>>('/messages',
        queryParameters: queryParameters,
        options: RequestOptions(
            method: 'GET',
            headers: <String, dynamic>{},
            extra: _extra,
            baseUrl: baseUrl),
        data: _data);
    var value = _result.data.map((k, dynamic v) => MapEntry(
        k,
        (v as List)
            .map((i) => Message.fromJson(i as Map<String, dynamic>))
            .toList()));

    return value;
  }

  @override
  Future<Message> postMessage(params) async {
    ArgumentError.checkNotNull(params, 'params');
    const _extra = <String, dynamic>{};
    final queryParameters = <String, dynamic>{};
    final _data = <String, dynamic>{};
    _data.addAll(params ?? <String, dynamic>{});
    final _result = await _dio.request<Map<String, dynamic>>('/message',
        queryParameters: queryParameters,
        options: RequestOptions(
            method: 'POST',
            headers: <String, dynamic>{},
            extra: _extra,
            baseUrl: baseUrl),
        data: _data);
    final value = Message.fromJson(_result.data);
    return value;
  }

  @override
  Future<Message> patchMessage(messageId, params) async {
    ArgumentError.checkNotNull(messageId, 'messageId');
    ArgumentError.checkNotNull(params, 'params');
    const _extra = <String, dynamic>{};
    final queryParameters = <String, dynamic>{};
    final _data = <String, dynamic>{};
    _data.addAll(params ?? <String, dynamic>{});
    final _result = await _dio.request<Map<String, dynamic>>(
        '/message/$messageId',
        queryParameters: queryParameters,
        options: RequestOptions(
            method: 'PATCH',
            headers: <String, dynamic>{},
            extra: _extra,
            baseUrl: baseUrl),
        data: _data);
    final value = Message.fromJson(_result.data);
    return value;
  }

  @override
  Future<void> deleteMessage(messageId) async {
    ArgumentError.checkNotNull(messageId, 'messageId');
    const _extra = <String, dynamic>{};
    final queryParameters = <String, dynamic>{};
    final _data = <String, dynamic>{};
    await _dio.request<void>('/message/$messageId',
        queryParameters: queryParameters,
        options: RequestOptions(
            method: 'DELETE',
            headers: <String, dynamic>{},
            extra: _extra,
            baseUrl: baseUrl),
        data: _data);
    return null;
  }
}
