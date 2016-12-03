package com.scheduler.record.radio.pidroidfmrec;

import android.content.pm.PackageInfo;
import android.content.pm.PackageManager;
import android.os.Bundle;
import android.support.annotation.NonNull;
import android.support.v7.app.AppCompatActivity;
import android.widget.TableLayout;
import android.widget.TableRow;
import android.widget.TextView;

import java.util.List;

public class ListInstalledAppsActivity extends AppCompatActivity
{

    private static final String TAG = "ListAppsActivity";

    @Override
    protected void onCreate(Bundle savedInstanceState)
    {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_list_installed_apps);
        TableLayout tableLayout = (TableLayout) findViewById(R.id.tableLayout_table);

        final PackageManager pm = getPackageManager();
        // Get a list of installed apps.
        List<PackageInfo> packages = pm.getInstalledPackages(PackageManager.GET_PERMISSIONS);

        assert tableLayout != null;
        for (PackageInfo packageInfo : packages)
        {
            TableRow tableRow = new TableRow(this);

            TextView textViewAppName = getTextView(pm.getApplicationLabel(packageInfo.applicationInfo).toString());
            TextView textViewPackageName = getTextView(packageInfo.packageName);

            tableRow.addView(textViewAppName);
            tableRow.addView(textViewPackageName);

            tableLayout.addView(tableRow);
        }
    }

    @NonNull
    private TextView getTextView(String text)
    {
        TextView textViewAppName = new TextView(this);
        textViewAppName.setText(text);
        return textViewAppName;
    }
}
