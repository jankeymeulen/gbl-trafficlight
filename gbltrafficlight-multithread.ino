#include <Arduino.h>
#include <WiFi.h>
#include <HTTPClient.h>
#include <FastLED.h>
#include <SPI.h>

#define NUM_LEDS 60
#define DATA_PIN 23
#define CLOCK_PIN 18

/* general LED stuff */
CRGB leds[NUM_LEDS];
CRGB colour = CRGB::White;
char effect[50] = "";
unsigned long previousFrame = 0;

/* blink stuff */
#define BLINK_DELAY 1000
bool blinkState = false;/* cylon stuff */
#define CYLON_WIDTH 20
#define CYLON_DELAY 50
int cylonStart = 0;
bool cylonDirection = true;

/* pong stuff */
#define PONG_WIDTH 5
#define PONG_MIN_DELAY 1
#define PONG_MAX_DELAY 20
#define PONG_CYCLE_LENGTH 2000
#define PONG_BLUR 172
int pongDelay = 20;
int pongCycleStart = 0;
int pongStart = 0;
bool pongDirection = true;
bool pongCycleDirection = true;

/* Run stuff */
#define RUN_WIDTH 10
#define RUN_DELAY 5
int runStart = 0;
bool runDirection = true;

/* Sparkle stuff */
#define SPARKLE_COUNT 1
#define SPARKLE_DIMM 244
#define SPARKLE_DELAY 20
#define SPARKLE_CYCLE 500
int previousSparkle = 0;

/* Rainbow stuff */
#define RAINBOW_DELAY 100
uint8_t rainbowHue = 0;

/* Theatre stuff */
#define THEATRE_DELAY 300
#define THEATRE_LENGTH 3
int theatreToggle = 0;

/* Breathe stuff */
#define BREATHE_DELAY 50
#define BREATHE_CYCLE 3000.0

/* Fadeblack stuff */
#define FADEBLACK_DELAY 50
#define FADEBLACK_CYCLE 3000

/* Fadewhite stuff */
#define FADEWHITE_DELAY 50
#define FADEWHITE_CYCLE 3000

/* Fire stuff */
#define FIRE_COOLING 55
#define FIRE_SPARKING 120
#define FIRE_DELAY 30
static byte heat[NUM_LEDS];
int cooldown;

/* Comet stuff */
#define COMET_WIDTH 10
#define COMET_DELAY 30
#define COMET_COOLING 55
#define COMET_INTERVAL 5000
int cometPosition = 0;
int nextComet = 0;
int cometToggle = false;

/* Strobo stuff */
#define STROBO_INTERVAL 100
#define STROBO_LENGTH 5
int nextStrobo = 0;
bool stroboToggle = false;

/* Lightning stuff */

/* Crash stuff */
#define CRASH_WIDTH 2

/* Belgium stuff */
#define BELGIUM_COUNT 3
#define BELGIUM_DELAY 40
int belgiumStart = 0;

/* France stuff */
#define FRANCE_COUNT 3
#define FRANCE_DELAY 40
int franceStart = 0;

void setup() {


  Serial.begin(115200);
  Serial.println();

  Serial.print("[SETUP] Connecting to Wifi");
  WiFi.begin("GoogleGuest-Legacy");
  long startConnect = millis();
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
    if (startConnect + 60000 < millis())
    {
      Serial.println("\nError connecting, rebooting");
      ESP.restart();
    }
  }
  Serial.println(" connected!");

  FastLED.addLeds<APA102, DATA_PIN, CLOCK_PIN, BGR, DATA_RATE_MHZ(1)>(leds, NUM_LEDS);
  FastLED.setBrightness(255);
  Serial.println("[SETUP] LEDs configured!");

  xTaskCreate(
    ledTask,          /* Task function. */
    "LedTask",        /* String with name of task. */
    10000,            /* Stack size in words. */
    NULL,             /* Parameter passed as input of the task */
    1,                /* Priority of the task. */
    NULL);            /* Task handle. */

  xTaskCreate(
    httpClientTask,          /* Task function. */
    "HttpClientTask",        /* String with name of task. */
    10000,            /* Stack size in words. */
    NULL,             /* Parameter passed as input of the task */
    1,                /* Priority of the task. */
    NULL);            /* Task handle. */

  Serial.println("[SETUP] Done!");

}

void loop() {
  delay(360000);
  ESP.restart();
}

void httpClientTask(void* parameter) {
  while (1) {
    HTTPClient http;
    if ((WiFi.status() == WL_CONNECTED)) {

      HTTPClient http;

      http.begin("http://gbl-trafficlight.appspot.com/effect"); //HTTP

      Serial.print("[HTTP] GET...\n");
      // start connection and send HTTP header
      int httpCode = http.GET();

      // httpCode will be negative on error
      if (httpCode > 0) {
        // HTTP header has been send and Server response header has been handled
        Serial.printf("[HTTP] GET done, code: %d\n", httpCode);

        // file found at server
        if (httpCode == HTTP_CODE_OK) {
          String payload = http.getString();
          char payloadchar[50];
          payload.toCharArray(payloadchar, 50);
          Serial.printf("[HTTP] Payload [%s]\n", payloadchar);
          int r, g, b;
          sscanf(payloadchar, "%s %d %d %d", effect, &r, &g, &b);
          colour.r = r;
          colour.g = g;
          colour.b = b;
          Serial.printf("[HTTP] Parsed Effect:[%s],R:[%d],G:[%d],B:[%d]\n", effect, r, g, b);
        }
      } else {
        Serial.printf("[HTTP] GET... failed, error: %s\n", http.errorToString(httpCode).c_str());
        ESP.restart();
      }

      http.end();
    }
    vTaskDelay(2500);
  }
}

void ledTask(void* parameter)
{
  Serial.println("[LED] Entered task");
  previousFrame = millis();
  while (1) {
    if ( strcmp(effect, "SOLID") == 0 ) {
      fillSolid();
    } else if ( strcmp(effect, "BLINK" ) == 0 ) {
      fillBlink();
    } else if ( strcmp(effect, "CYLON" ) == 0 ) {
      fillCylon();
    } else if ( strcmp(effect, "RUN" ) == 0 ) {
      fillRun();
    } else if ( strcmp(effect, "PONG" ) == 0 ) {
      fillPong();
    } else if ( strcmp(effect, "SPARKLE" ) == 0 ) {
      fillSparkle();
    } else if ( strcmp(effect, "RAINBOW" ) == 0 ) {
      fillRainbow();
    } else if ( strcmp(effect, "THEATRE" ) == 0 ) {
      fillTheatre();
    } else if ( strcmp(effect, "FIRE" ) == 0 ) {
      fillFire();
    } else if ( strcmp(effect, "COMET" ) == 0 ) {
      fillComet();
    } else if ( strcmp(effect, "BREATHE" ) == 0 ) {
      fillBreathe();
    } else if ( strcmp(effect, "STROBO" ) == 0 ) {
      fillStrobo();
    } else if ( strcmp(effect, "BELGIUM" ) == 0 ) {
      fillBelgium();
    } else if ( strcmp(effect, "FRANCE" ) == 0 ) {
      fillFrance();
    }
    FastLED.show();
    vTaskDelay(1);
  }
}

void fillBlink() {
  if (previousFrame + BLINK_DELAY < millis()) {
    if (blinkState) {
      fill_solid (leds, NUM_LEDS, colour);
    } else {
      fill_solid (leds, NUM_LEDS, CRGB::Black);
    }
    blinkState = !blinkState;
    previousFrame = millis();
  }
}

void fillStrobo() {
  if (nextStrobo < millis()) {
    if (stroboToggle) {
      fill_solid (leds, NUM_LEDS, colour);
      nextStrobo = millis() + STROBO_LENGTH;
    } else {
      fill_solid (leds, NUM_LEDS, CRGB::Black);
      nextStrobo = millis() + STROBO_INTERVAL;
    }
    stroboToggle = !stroboToggle;
  }
}

void fillBelgium() {
  if (previousFrame + BELGIUM_DELAY < millis()) {
    int flagWidth = ceil(NUM_LEDS/BELGIUM_COUNT);
    
    for (int i = 0; i < NUM_LEDS; i++)
    {
      int segment = floor(((i+belgiumStart)%flagWidth)/ceil(flagWidth/3));
      if (segment == 0) {
        leds[i] = CRGB::Red;
      } else if (segment == 1) {
        leds[i] = CRGB::Yellow;
      } else {
        leds[i] = CRGB::Black;
      }
    }
    belgiumStart++;
    if(belgiumStart >= NUM_LEDS)
    {
      belgiumStart = 0;
    }

    previousFrame = millis();
  }
}

void fillFrance() {
  if (previousFrame + FRANCE_DELAY < millis()) {
    int flagWidth = ceil(NUM_LEDS/FRANCE_COUNT);
    
    for (int i = 0; i < NUM_LEDS; i++)
    {
      int segment = floor(((i+franceStart)%flagWidth)/ceil(flagWidth/3));
      if (segment == 0) {
        leds[i] = CRGB::Red;
      } else if (segment == 1) {
        leds[i] = CRGB::White;
      } else {
        leds[i] = CRGB::Blue;
      }
    }
    franceStart++;
    if(franceStart >= NUM_LEDS)
    {
      franceStart = 0;
    }

    previousFrame = millis();
  }
}

void fillCylon() {
  if (previousFrame + CYLON_DELAY < millis()) {
    fill_solid(leds, NUM_LEDS, CRGB::Black);
    for (int i = cylonStart; i < NUM_LEDS && i < cylonStart + CYLON_WIDTH ; i++) {
      leds[i] = colour;
    }

    if ( cylonDirection ) {
      cylonStart++;
    } else {
      cylonStart--;
    }

    if (cylonStart >= NUM_LEDS - CYLON_WIDTH) {
      cylonDirection = false;
    }

    if (cylonStart == 0) {
      cylonDirection = true;
    }
    previousFrame = millis();
  }
}

void fillPong() {
  if (previousFrame + pongDelay < millis()) {
    fill_solid(leds, NUM_LEDS, CRGB::Black);
    for (int i = pongStart; i < NUM_LEDS && i < pongStart + PONG_WIDTH ; i++) {
      leds[i] = colour;
    }
    blur1d(leds, NUM_LEDS, PONG_BLUR);

    if ( pongDirection ) {
      pongStart++;
    } else {
      pongStart--;
    }

    if (pongStart >= NUM_LEDS - PONG_WIDTH) {
      pongDirection = false;
    }

    if (pongStart == 0) {
      pongDirection = true;
    }

    if (pongCycleStart + PONG_CYCLE_LENGTH < millis() )
    {
      if ( pongCycleDirection ) {
        pongDelay--;
      } else {
        pongDelay++;
      }
      Serial.printf("[DEBUG] Pong Cycled, delay [%d]\n", pongDelay);
      pongCycleStart = millis();
    }

    if (pongDelay >= PONG_MAX_DELAY) {
      pongCycleDirection = true;
    }

    if (pongDelay <= PONG_MIN_DELAY) {
      pongCycleDirection = false;
    }
    previousFrame = millis();
  }
}

void fillRun() {
  if (previousFrame + RUN_DELAY < millis()) {
    fill_solid(leds, NUM_LEDS, CRGB::Black);
    for (int i = runStart; i < NUM_LEDS && i < runStart + RUN_WIDTH ; i++) {
      leds[i] = colour;
    }
    if (runStart >= NUM_LEDS - RUN_WIDTH)
      runStart = 0;
    runStart++;
    previousFrame = millis();
  }
}

//void fillComet2() {
//  if (previousFrame + COMET_DELAY < millis()) {
//    fill_solid(leds, NUM_LEDS, CRGB::Black);
//    for (int i = cometStart; i >= 0 && i > cometStart - COMET_WIDTH ; i--) {
//      leds[i] = colour;
//      leds[i].fadeLightBy(((cometStart-i) * 256) / COMET_WIDTH);
//    }
//    if (cometStart >= NUM_LEDS)
//      cometStart = 0;
//    cometStart++;
//    previousFrame = millis();
//  }
//}

void fillSparkle() {
  if (previousFrame + SPARKLE_DELAY < millis()) {

    for (int i = 0; i < NUM_LEDS; i++) {
      leds[i].nscale8( SPARKLE_DIMM );
    }

    if ( previousSparkle + SPARKLE_CYCLE < millis() ) {
      for (int i = 0; i < SPARKLE_COUNT; i++) {
        leds[random(NUM_LEDS)] = colour;
      }
    }
    previousFrame = millis();
  }
}

void fillRainbow() {
  if (previousFrame + RAINBOW_DELAY < millis()) {
    fill_rainbow(leds, NUM_LEDS, rainbowHue);
    rainbowHue--;
    previousFrame = millis();
  }
}

void fillSolid() {
  fill_solid (leds, NUM_LEDS, colour);
}

void fillBreathe() {
  if ( previousFrame + BREATHE_DELAY < millis() ) {
    previousFrame = millis();
    CRGB dimmedColour = colour;
    dimmedColour %= (exp(sin(previousFrame / BREATHE_CYCLE * PI)) - 0.36787944) * 108.0;
    dimmedColour += 10;
    fill_solid (leds, NUM_LEDS, dimmedColour);
  }
}

void fillFadeBlack() {
  if ( previousFrame + FADEBLACK_DELAY < millis() ) {
    previousFrame = millis();
    CRGB dimmedColour = colour;
    dimmedColour %= (exp(sin(previousFrame / BREATHE_CYCLE * PI)) - 0.36787944) * 108.0;
    dimmedColour += 10;
    fill_solid (leds, NUM_LEDS, dimmedColour);
  }
}


void fillTheatre() {
  if ( previousFrame + THEATRE_DELAY < millis() ) {
    for ( int i = 0; i < NUM_LEDS; i++ )
    {
      if (i % THEATRE_LENGTH == theatreToggle)
      {
        leds[i] = colour;
      } else {
        leds[i] = CRGB::Black;
      }
    }
    theatreToggle++;
    if (theatreToggle == THEATRE_LENGTH)
    {
      theatreToggle = 0;
    }
    previousFrame = millis();
  }
}

void fillComet() {

  if ( previousFrame + COMET_DELAY < millis() ) {

    // Cooldown all cells
    for ( int i = 0; i < NUM_LEDS; i++) {
      cooldown = random(0, ((COMET_COOLING * 10) / NUM_LEDS) + 2);

      if (cooldown > heat[i]) {
        heat[i] = 0;
      } else {
        heat[i] = heat[i] - cooldown;
      }
    }

    // Move the comet up
    if (cometPosition < NUM_LEDS - 1 && cometToggle)
    {
      heat[cometPosition] = 255;
      heat[cometPosition + 1] = 255;
      cometPosition += 2;
    }

    // Convert heat to colours
    for ( int j = 0; j < NUM_LEDS; j++) {
      setPixelHeatColor(j, heat[j] );
    }

    // Colour the comet to the custom colour
    if (cometPosition < NUM_LEDS - 1 && cometToggle)
    {
      leds[cometPosition + 1] = colour;
    }

    if (cometPosition >= NUM_LEDS)
    {
      cometToggle = false;
      cometPosition = 0;
      nextComet = millis() + random(COMET_INTERVAL, COMET_INTERVAL * 3);
    }

    if (nextComet < millis())
    {
      cometToggle = true;
    }

    previousFrame = millis();
  }
}

void fillFire() {

  if ( previousFrame + FIRE_DELAY < millis() ) {

    // Step 1.  Cool down every cell a little
    for ( int i = 0; i < NUM_LEDS; i++) {
      cooldown = random(0, ((FIRE_COOLING * 10) / NUM_LEDS) + 2);

      if (cooldown > heat[i]) {
        heat[i] = 0;
      } else {
        heat[i] = heat[i] - cooldown;
      }
    }

    // Step 2.  Heat from each cell drifts 'up' and diffuses a little
    for ( int k = NUM_LEDS - 1; k >= 2; k--) {
      heat[k] = (heat[k - 1] + heat[k - 2] + heat[k - 2]) / 3;
    }

    // Step 3.  Randomly ignite new 'sparks' near the bottom
    if ( random(255) < FIRE_SPARKING ) {
      int y = random(7);
      heat[y] = heat[y] + random(160, 255);
    }

    // Step 4.  Convert heat to LED colors
    for ( int j = 0; j < NUM_LEDS; j++) {
      setPixelHeatColor(j, heat[j] );
    }
    previousFrame = millis();
  }
}

void setPixelHeatColor (int pixel, byte temperature) {
  // Scale 'heat' down from 0-255 to 0-191
  byte t192 = round((temperature / 255.0) * 191);

  // calculate ramp up from
  byte heatramp = t192 & 0x3F; // 0..63
  heatramp <<= 2; // scale up to 0..252

  // figure out which third of the spectrum we're in:
  if ( t192 > 0x80) {                    // hottest
    leds[pixel].setRGB(255, 255, heatramp);
  } else if ( t192 > 0x40 ) {            // middle
    leds[pixel].setRGB(255, heatramp, 0);
  } else {                               // coolest
    leds[pixel].setRGB(heatramp, 0, 0);
  }
}


