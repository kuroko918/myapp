import 'package:dio/dio.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:pretty_dio_logger/pretty_dio_logger.dart';
import 'package:retrofit/retrofit.dart';
import 'package:myapp/utils/firebase/firebase_auth.dart';
import 'package:myapp/domains/user.dart';

part 'user_api_client.g.dart';

abstract class UserApiClient {
  factory UserApiClient(Dio dio, {String baseUrl}) = _UserApiClient;

  static Future<UserApiClient> create() async {
    final dio = Dio();
    final idToken = await getIdToken();
    final baseUrl = env['BASE_URL'];
    dio.options.headers['Authorization'] = 'bearer $idToken';
    dio.interceptors.add(PrettyDioLogger());
    return UserApiClient(dio, baseUrl: baseUrl);
  }

  @GET('/user/{id}')
  Future<User> getUser(@Path('id') String userId);

  @PUT('/user/{id}')
  Future<User> putUser(@Path('id') String userId, @Body() User user);
}
