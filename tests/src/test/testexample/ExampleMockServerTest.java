package testexample;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;

import org.junit.Assert;
import org.junit.BeforeClass;
import org.junit.Test;

public class ExampleMockServerTest {

  @BeforeClass
  public static void setup() {
    ExampleMockServer server = new ExampleMockServer();
    server.createExpectationMockServerClient();
  }

  @Test
  public void testMockServer() throws IOException {

    // make a call to /view/cart
    URL url = new URL("http://127.0.0.1:8001/view/cart");
    HttpURLConnection con = (HttpURLConnection)url.openConnection();
    con.setRequestMethod("GET");
    int responseCode = con.getResponseCode();

    BufferedReader in = new BufferedReader(
        new InputStreamReader(con.getInputStream()));
    String inputLine;
    StringBuffer responseContent = new StringBuffer();
    while ((inputLine = in.readLine()) != null) {
      responseContent.append(inputLine);
    }
    in.close();
    String responseContentString = responseContent.toString();

    // then
    Assert.assertEquals(200, responseCode);
    Assert.assertEquals("some_response_body", responseContentString);
  }
}
