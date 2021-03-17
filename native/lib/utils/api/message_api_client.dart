import 'package:dio/dio.dart';
import 'package:pretty_dio_logger/pretty_dio_logger.dart';
import 'package:retrofit/retrofit.dart';
import 'package:myapp/utils/firebase/firebase_auth.dart';
import 'package:myapp/domains/message.dart';

part 'message_api_client.g.dart';

@RestApi(baseUrl: 'http://localhost:8080')
abstract class MessageApiClient {
  factory MessageApiClient(Dio dio, {String baseUrl}) = _MessageApiClient;

  static Future<MessageApiClient> create() async {
    final dio = Dio();
    final idToken = await getIdToken();
    dio.options.headers['Authorization'] = 'bearer $idToken';
    dio.interceptors.add(PrettyDioLogger());
    return MessageApiClient(dio);
  }

  @GET('/messages')
  Future<Map<String, List<Message>>> getMessages();

  @POST('/message')
  Future<Message> postMessage(@Body() Map<String, dynamic> params);

  @PATCH('/message/{id}')
  Future<Message> patchMessage(@Path('id') String messageId, @Body() Map<String, dynamic> params);

  @DELETE('/message/{id}')
  Future<void> deleteMessage(@Path('id') String messageId);
}
