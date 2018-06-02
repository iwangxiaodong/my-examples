

https://webrtc.org/native-code/android/


https://webrtc.googlesource.com/src/+/master/examples/androidapp


git clone https://webrtc.googlesource.com/src


【 or (60MB)】 wget https://webrtc.googlesource.com/src/+archive/master.tar.gz


cd src-master/examples/androidapp

wget https://raw.githubusercontent.com/iwangxiaodong/my-examples/master/webrtc/apprtc/build.gradle


use android studio open or direct gradle build


注意：为避免显式申请权限暂用targetSdkVersion 22


    Android应用若报Handshake failed，可暂时信任所有证书： org.appspot.apprtc.util.AsyncHttpURLConnection.sendHttpMessage()
        
        if (connection instanceof HttpsURLConnection) {
            HttpsURLConnection httpsConn = (HttpsURLConnection) connection;
            httpsConn.setSSLSocketFactory(SSLCertificateSocketFactory.getInsecure(0, null));
            httpsConn.setHostnameVerifier(new AllowAllHostnameVerifier());
        }
