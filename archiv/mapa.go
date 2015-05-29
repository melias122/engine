package archiv

/*
void mapaNtice(CSV csv, uint n){

    CSV mapaNtice;
    QStringList out_stringl, header, ntica_anl;
    Kombinacia akt;
    QHash<QString, int> ntice_pos, ntice_krok;

    // ntice
    ntice_pos = Ntice(n);
    ntice_krok = QHash<QString, int>(ntice_pos);
    ntice_pocet = QHash<QString, int>(ntice_pos);
    foreach (const QString &strp, ntice_krok.keys()) {
        ntice_krok[strp] = -1;
        ntice_pocet[strp] = 0;
    }

    // header
    header << "N-tica";
    for(uint i=1; i <= n; i++)
        header << QString::number(i);
    header << "Sucet N-tic" << "Sucin pozicie a stlpca";

    for(int i=0; i < ntice_pos.size(); i++)
        header << QString(ntice_pos.key(i)).prepend("Krok ");

    out_stringl << csv.front() << header;
    mapaNtice.append(out_stringl);

    csv.pop_front();

    // analyza
    int pos=0;
    foreach (const QStringList &akt_r, csv) {
        pos++;

        if(!akt.isEmpty())
            akt.clear();

        for(uint i=3; i < n+3;i++)
            akt.push_back(akt_r[i].toInt());

        if(!ntica_anl.empty())
            ntica_anl.clear();

        for(uint i=0; i < n+3+ntice_pos.size(); i++){
            ntica_anl << "";
        }
        ntica_anl[0] = nticaToQString(ntica(akt));

        if(ntica_anl[0] != ntice_pos.key(0)){ // ak neni n 0 0 0 ...

            int start=-1, end=-1;
            for(int i=0; i<akt.size()-1;i++){

                for(int j=i; j <akt.size()-1;j++){
                    if((akt[j+1] - akt[j]) == 1){
                        if(start==-1)
                            start=j;
                        end=j+1;
                        i=j;
                    }
                    else
                        break;
                }
                if(start != end) {
                    int sum=0,sum_stl=1;
                    for(int j=start; j <= end; j++){
                        sum += (int)akt[j];
                        sum_stl *= (j+1);
                        // cislo v stl
                        ntica_anl[j+1] = QString::number(akt[j]);
                    }

                    // sucet ntic
                    if(ntica_anl[n+1].size() > 0)
                        ntica_anl[n+1].append(QString::number(sum).prepend(", "));
                    else
                        ntica_anl[n+1] = QString::number(sum);

                    // sucin stlpcov
                    if(ntica_anl[n+2].size() > 0)
                        ntica_anl[n+2].append(QString::number(sum_stl).prepend(", "));
                    else
                        ntica_anl[n+2] = QString::number(sum_stl);

                    start=end=-1;
                }
            }
        }

        // ntice stat
        QString tica = ntica_anl[0], sucin_stl = ntica_anl[n+2];
        ntice_pocet[tica]++;
        if(!sucin_stl.isEmpty()){
            if(ntice_typ_pocet[tica].value(sucin_stl) > 0){
                int p = ntice_typ_pocet[tica].value(sucin_stl);
                ntice_typ_pocet[tica].insert(sucin_stl, p + 1);
            }
            else{
                ntice_typ_pocet[ntica_anl[0]].insert(ntica_anl[n+2], 1);
            }
        }
        //

        foreach (const QString &strp, ntice_krok.keys()) {
            if(ntice_krok[strp] != -1)
                ntice_krok[strp]++;
        }

        if(ntice_krok[ntica_anl[0]] == -1){
            ntica_anl[n+3+ntice_pos[ntica_anl[0]]] = QString::number(0);
        }
        else{
            ntica_anl[n+3+ntice_pos[ntica_anl[0]]] = QString::number(ntice_krok[ntica_anl[0]]);
        }
        ntice_krok[ntica_anl[0]] = 0;


        out_stringl.clear();
        out_stringl << akt_r << ntica_anl;

        mapaNtice.append(out_stringl);
    }
    exportCsv(mapaNtice, pwd() + "MapaNtice_" + suborName() + ".csv");
}

void mapaZhoda(CSV csv, uint n){

    CSV mapaZhoda;
    int zhoda{0}, l_len{0}, krok{-1};
    QStringList out_stringl, header, zh_anl;
    QVector<int> prd ,akt;
    QHash<int, int> zh_pos;

    // dlzka riadku
    if(csv.size() > 0)
        l_len = csv[1].length();
    else
        l_len = csv.front().length();

    // header
    header << "Zhoda";
    for(uint i{1}; i <= n; ++i)
        header << QString::number(i);
    header << "Pozicia ZH" << "Krok";
    for(uint i{1}; i <= n; ++i)
        header << QString::number(i).append(" Stl r-1/r");

    out_stringl << csv.front() << header;
    mapaZhoda.append(out_stringl);
    // header end

    //header pop
    csv.pop_front();

    for(uint i{0}; i <= n; ++i)
        zhoda_pocet.insert(i, 0);

    foreach (const QStringList &akt_r, csv){

        if(!akt.isEmpty())
            akt.clear();

        for(uint i=3; i < n+3;i++)
            akt.push_back(akt_r[i].toInt());

        zhoda=0;
        zh_anl.clear();
        zh_pos.clear();

        for(uint i=0; i < 2*n+3;i++){
            zh_anl << "";
        }

        for(int i=0; i<akt.size() && !prd.empty(); i++){
            for(int j=0; j <prd.size();j++){
                if(akt[i] == prd[j]){
                    zhoda++;
                    zh_pos.insert(i,j);
                }
            }
        }

        // Zhoda
        zh_anl[0] = QString::number(zhoda);

        // Krok
        if(zhoda > 0){
            if(krok == -1)
                zh_anl[n+2] = QString::number(0);
            else
                zh_anl[n+2] = QString::number(krok);
            krok=1;
        }
        else if(krok >= 0)
            krok++;

        QString pos_zh{""};

        auto zhPosKeys = zh_pos.keys();
        if(!zhPosKeys.isEmpty())
            qSort(zhPosKeys);

        for(int &i : zhPosKeys){
            // 1-n
            mapaZhoda.back()[l_len+1+zh_pos[i]] = QString::number(prd[zh_pos[i]]);
            zh_anl[i+1] = QString::number(akt[i]);

            // Pozicia zhoda
            pos_zh = QString::number(zh_pos[i]+1).append("|").append(QString::number(i+1));
            if(zh_anl[n+1].size() > 0){
//                pos_zh += ", " + QString::number(zh_pos[i]+1).append("|").append(QString::number(i+1));
                zh_anl[n+1].append(", ").append(pos_zh);
            }
            else{
//                pos_zh = QString::number(zh_pos[i]+1).append("|").append(QString::number(i+1));
                zh_anl[n+1].append(pos_zh);
            }

            // nstl r-1/r
            zh_anl[n+3+i] = pos_zh;
        }

        //stat aln
//        auto headerTmp = mapaZhoda.takeFirst();
//        for(auto &riadok : mapaZhoda){

//            QString zhs = riadok[2*n+4];
        QString zhs = zh_anl[n+1];

//        qDebug() << zhs << zh_anl;

        int zhoda_kolko;
        if(!zhs.isEmpty())
            zhoda_kolko = zhs.split(",").size();
        else
            zhoda_kolko = 0;
//        qDebug() << zhoda_kolko << zhs;
        zhoda_pocet[zhoda_kolko]++;

        if(zhoda_kolko > 0){  // len zhoda 1,2,3,..
            if(!zhoda_typ_pocet[zhoda_kolko].contains(zhs))
                zhoda_typ_pocet[zhoda_kolko].insert(zhs,1);
            else
                zhoda_typ_pocet[zhoda_kolko][zhs]++;
        }
//        }
//        mapaZhoda.prepend(headerTmp);
        //

        out_stringl.clear();
        out_stringl << akt_r << zh_anl;

        mapaZhoda.push_back(QStringList(out_stringl));

        prd = akt;
    }

//    //stat aln
//    auto headerTmp = mapaZhoda.takeFirst();
//    for(auto &riadok : mapaZhoda){

//        QString zhs = riadok[2*n+4];
//        int zhoda_kolko;
//        if(!zhs.isEmpty())
//            zhoda_kolko = zhs.split(",").size();
//        else
//            zhoda_kolko = 0;
//        qDebug() << zhoda_kolko << zhs;
//        zhoda_pocet[zhoda_kolko]++;

//        if(zhoda_kolko > 0){  // len zhoda 1,2,3,..
//            if(!zhoda_typ_pocet[zhoda_kolko].contains(zhs))
//                zhoda_typ_pocet[zhoda_kolko].insert(zhs,1);
//            else
//                zhoda_typ_pocet[zhoda_kolko][zhs]++;
//        }
//    }
//    mapaZhoda.prepend(headerTmp);
//    //

    exportCsv(mapaZhoda, pwd() + "MapaZhoda_" + suborName() + ".csv");
}

*/

// func (a *Archiv) MapaXtice(f *os.File) error {
// 	f.Seek(0, 0)

// 	r := csv.NewReader(f)
// 	r.Comma = ';'
// 	rows, err := r.ReadAll()
// 	if err != nil {
// 		return err
// 	}

// 	file := xlsx.NewFile()
// 	sheet := file.AddSheet("Mapa Xtice")

// 	for i, r := range rows {
// 		row := sheet.AddRow()
// 		for j, c := range r {
// 			cell := row.AddCell()

// 			cell.Value = c

// TODO: farby
// style := xlsx.NewStyle()
// switch (c - 1) / 10 {
// case 0: // Red
// 	style.Fill = *xlsx.NewFill("solid", "A0000000", "FFFFAAAA")
// case 1: // Green
// 	style.Fill = *xlsx.NewFill("solid", "BFFFAAAA", "AA000000")
// case 2: // Blue
// 	style.Fill = *xlsx.NewFill("solid", "CFFFAAAA", "BB000000")
// case 3: // Yellow
// 	style.Fill = *xlsx.NewFill("solid", "DFFFAAAA", "CC000000")
// case 4: // Cyan
// 	style.Fill = *xlsx.NewFill("solid", "EFFFAAAA", "DD000000")
// case 5: // Dark Red
// 	style.Fill = *xlsx.NewFill("solid", "FFFFAAAA", "EE000000")
// case 6: // Dark Green
// 	style.Fill = *xlsx.NewFill("solid", "AAFFAAAA", "FF000000")
// case 7: // Dark Blue
// 	style.Fill = *xlsx.NewFill("solid", "BBFFAAAA", "AAA00000")
// case 8: // Dark Yellow
// 	style.Fill = *xlsx.NewFill("solid", "CCFFAAAA", "BBB00000")
// default:
// }
// cell.SetStyle(style)
// }
// }

// 	return file.Save(fmt.Sprintf("%d%d/MapaXtice_%d%d.xlsx", a.n, a.m, a.n, a.m))
// }
