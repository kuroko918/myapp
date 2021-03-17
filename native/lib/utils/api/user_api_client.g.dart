// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'user_api_client.dart';

// **************************************************************************
// RetrofitGenerator
// **************************************************************************

class _UserApiClient implements UserApiClient {
  _UserApiClient(this._dio, {this.baseUrl}) {
    ArgumentError.checkNotNull(_dio, '_dio');
    baseUrl ??= 'http://localhost:8080';
  }

  final Dio _dio;

  String baseUrl;

  @override
  Future<User> getUser(userId) async {
    ArgumentError.checkNotNull(userId, 'userId');
    const _extra = <String, dynamic>{};
    final queryParameters = <String, dynamic>{};
    final _data = <String, dynamic>{};
    final _result = await _dio.request<Map<String, dynamic>>('/user/$userId',
        queryParameters: queryParameters,
        options: RequestOptions(
            method: 'GET',
            headers: <String, dynamic>{},
            extra: _extra,
            baseUrl: baseUrl),
        data: _data);
    final value = User.fromJson(_result.data);
    return value;
  }

  @override
  Future<User> putUser(userId, user) async {
    ArgumentError.checkNotNull(userId, 'userId');
    ArgumentError.checkNotNull(user, 'user');
    const _extra = <String, dynamic>{};
    final queryParameters = <String, dynamic>{};
    final _data = <String, dynamic>{};
    _data.addAll(user?.toJson() ?? <String, dynamic>{});
    final _result = await _dio.request<Map<String, dynamic>>('/user/$userId',
        queryParameters: queryParameters,
        options: RequestOptions(
            method: 'PUT',
            headers: <String, dynamic>{},
            extra: _extra,
            baseUrl: baseUrl),
        data: _data);
    final value = User.fromJson(_result.data);
    return value;
  }
}
