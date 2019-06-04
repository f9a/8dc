---
title: Setup Gitlab-CI for mobile
date: 2019-04-23T00:22:09+01:00
categories:
  - gitlab
  - ci
  - mobile
  - ios
  - android
url: /post/uuid/1af732d8-5bde-589c-a282-db05528e2d20
---

# Copy CI files from cimonster Project

https://gitlab.ants.house/cimonster/fe/mobile

```bash
cp -a ~/cimonster/fe/mobile/fastlane .
cp ~/cimonster/fe/mobile/Gemfile .
bundle update
cp ~/cimonster/fe/mobile/.gitlab-ci.yml .
cp -a ~/cimonster/fe/mobile/ci .
```

# gitlab-ci.yml

Wenn Projekt nicht von der standard Struktur abweicht muss nur die PROJECT_NAME Variable angepasst werden.

# iOS

## Certificates, Identifiers & Profiles

https://developer.apple.com/account/resources/certificates/list

Create new App ID

Create new distribution profile $projectname-DISTRIBUTION.

Create new development profile $projectname-DEV.

## Appstoreconnect

[https://appstoreconnect.apple.com](https://appstoreconnect.apple.com/) 

Create app

## Xcode

Setze Bundle Identifier auf house.ants.$projectname.

Deaktiviert "Automatically manage signing".

Erstelle eine neues Scheme in dem folgeden Format $projectname-RELEASE. Setzte Configuration auf Release.

## Fastline match

Add provisioning profile to match cert repository (https://gitlab.ants.house/op/fastlane-match-add-cert)

```bash
ruby main.rb -I -r https://gitlab.ants.house/certs/app-store -u $admin_mail -t U8L56RYFUT
```

## Set gitlab-ci Envs

Recommendation choose password without symbols can happen that it you got problems with shell (bash/zsh/etc)

FASTLANE_PASSWORD is password for ci@ants.house apple-id

MATCH_PASSWORD is password for match certificate repository

# Android

## Play Store

Create App, setup basics.

If not existing create google credentials https://docs.fastlane.tools/getting-started/android/setup/#collect-your-google-credentials.

Store JSON in GOOGLE_PLAYSTORE_CREDENTIALS gitlab environment variable 

## keystore

**DON'T STORE KEYSTORE IN GIT REPO**

```bash
keytool -genkeypair -v -keystore $projectname.keystore -alias $projectname -keyalg RSA -keysize 2048 -validity 10000
keytool -importkeystore -srckeystore $projectname.keystore -destkeystore $projectname.keystore -deststoretype pkcs12
cat $projectname.keystore | base64
```

Store output in KEYSTORE gitlabe environment variable.

## Update gradle files

Update android/app/build.gradle

```gradle
    ....
    defaultConfig {
        applicationId "house.ants.$projectname"
        ....
```

### Enable signing

https://facebook.github.io/react-native/docs/signed-apk-android (we use different environment variables then on the webpage but the idea is the same)

```gradle
...
android {
    ...
    defaultConfig { ... }
    signingConfigs {
        release {
            if (project.hasProperty('APP_RELEASE_STORE_FILE')) {
                storeFile file(APP_RELEASE_STORE_FILE)
                storePassword APP_RELEASE_STORE_PASSWORD
                keyAlias APP_RELEASE_KEY_ALIAS
                keyPassword APP_RELEASE_KEY_PASSWORD
            }
        }
    }
    buildTypes {
        release {
            ...
            signingConfig signingConfigs.release
        }
    }
}
...
```

### Gradle properties

Update android/gradle.properties

```gradle
APP_RELEASE_STORE_FILE=$projectname.keystore
APP_RELEASE_KEY_ALIAS=$projectname  
```

### Set gitlab-ci Environment Variables

ORG_GRADLE_PROJECT_APP_RELEASE_STORE_PASSWORD password for keystore

ORG_GRADLE_PROJECT_APP_RELEASE_KEY_PASSWORD password for keystore alias item

# Enable Slack notifications (Optional)

Create a webhook for $projectname channel and add it to SLACK_URL gitlab-ci environment variable