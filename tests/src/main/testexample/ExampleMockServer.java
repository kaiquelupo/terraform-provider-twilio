package testexample;

import static org.mockserver.model.HttpRequest.request;
import static org.mockserver.model.HttpResponse.response;

import org.mockserver.client.MockServerClient;

public class ExampleMockServer {

  public void createExpectationMockServerClient() {
    new MockServerClient("localhost", 8001)
        .when(
            request()
                .withMethod("GET")
                .withPath("/view/cart")
//                .withCookies(
//                    cookie("session", "4930456C-C718-476F-971F-CB8E047AB349")
//                )
//                .withQueryStringParameters(
//                    param("cartId", "055CA455-1DF7-45BB-8535-4F83E7266092")
//                )
        )
        .respond(
            response()
                .withBody("some_response_body")
        );
  }
}
