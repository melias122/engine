package archiv

// void statistikaZhoda(uint n, uint m, uint csvLength){

//     CSV statZH;
//     QStringList out_stringl;
//     bigFloat cr;

//     // header
//     out_stringl << "Zhoda" << "Pocetnost teor." << "Teoreticka moznost v %" << "Pocetnost" << "Realne dosiahnute %";
//     statZH.append(out_stringl);

//     for(int i = n; i >= 0; --i){

//         out_stringl.clear();

//         cr = nCm(i,n).convert_to<bigFloat>() * nCm(m-n-(n-i),m-n).convert_to<bigFloat>();
//         out_stringl << QString::number(i)                       // ZH
//                     << doubleToQString(cr.convert_to<double>()); // pocet teor

//         cr /= nCm(n, m).convert_to<bigFloat>();
//         out_stringl << doubleToQString(cr.convert_to<double>()*100)         // teor %
//                     << QString::number(zhoda_pocet[i]);    // pocet

//         out_stringl << doubleToQString(((double)zhoda_pocet[i]/csvLength)*100);     // real %

//         statZH.append(out_stringl);
//     }
//     statZH.append({";"});

//     for(uint i=1; i<=n; i++){

//         out_stringl.clear();

//         QString zhs = QString::number(i);

//         // header
//         out_stringl << QString("Zhoda ").append(zhs) << "Pocetnost" << QString("Realne %");
//         statZH.append(out_stringl);
//         out_stringl.clear();
//         //

//         // vsetky
//         out_stringl << QString("Zhoda ").append(zhs) << QString::number(zhoda_pocet[i]) << doubleToQString(((double)zhoda_pocet[i]/csvLength)*100);
//         statZH.append(out_stringl);
//         //

//         foreach (const QString &poz, zhoda_typ_pocet[i].keys()) {
//             out_stringl.clear();
//             out_stringl << poz // poz
//                         << QString::number(zhoda_typ_pocet[i][poz]) // pocet
//                         << doubleToQString(((double)zhoda_typ_pocet[i][poz]/csvLength)*100);
//             statZH.append(out_stringl);
//         }
//         statZH.append({";"});
//     }

//     zhoda_pocet.clear();
//     zhoda_typ_pocet.clear();

//     zhoda_pocet.squeeze();
//     zhoda_typ_pocet.squeeze();

//     exportCsv(statZH, pwd() + "StatistikaZhoda_" + suborName() + ".csv");
// }

// void statistikaNtice(uint n, uint m, uint csvLength){

//     CSV statNtice;
//     QStringList out_stringl;
//     QHash<QString, int> ntice_pos;
//     QHash<QString, bigInt> ntice_poc_teor;

//     const QStringList header{"N-tica","Pocetnost teor.", "Teoreticka moznost v %", "Realne dosiahnute %"};
//     statNtice.append(header);

//     // ntice
//     ntice_pos = Ntice(n);

//     foreach (const QString &ntica, ntice_pos.keys()) {

//         int tica, k = m-n+1;
//         bigInt poc_real{1};
//         QVector<QString> qs = ntica.split(" ").toVector();

//         for(int i=0; i<qs.size(); i++){
//             if(qs[i].toInt() == 0)
//                 continue;
//             tica = qs[i].toInt();
//             poc_real *= nCm(k-tica,k);
//             k -= tica;
//         }
//         ntice_poc_teor.insert(ntica, poc_real);
//     }

//     for(int i=ntice_pos.size()-1; i>=0; i--){

//         QString tica = ntice_pos.key(i);
//         bigFloat cr;

//         out_stringl << tica << ntice_poc_teor.value(tica).str().c_str();

//         //tero moznost
//         cr.assign(ntice_poc_teor.value(tica).convert_to<bigFloat>()/nCm(n,m).convert_to<bigFloat>());
//         out_stringl << doubleToQString(cr.convert_to<double>()*100);

//         //real dosiah
//         out_stringl << doubleToQString((ntice_pocet[tica]/(double)csvLength)*100);

//         statNtice.append(out_stringl);
//         out_stringl.clear();
//     }

//     statNtice.append({";"});
//     out_stringl.clear();
//     out_stringl << "N-tica;" << "Sucin pozicie a stlpca;" << "Pocet vyskytov;" << "%" << "\n";
//     statNtice.append(out_stringl);

//     for(int i=ntice_pos.size()-1; i>=0; i--){
//         QString tica = ntice_pos.key(i);
//         QHash<QString, int> qmi = ntice_typ_pocet.value(tica);

//         out_stringl.clear();
//         out_stringl << tica << "vsetky" << QString::number(ntice_pocet.value(tica)) << doubleToQString(ntice_pocet[tica]/(double)csvLength*100);
//         statNtice.append(out_stringl);

//         foreach (const QString &p, qmi.keys()) {
//             out_stringl.clear();
//             out_stringl << tica << p << QString::number(qmi.value(p)) << doubleToQString((qmi.value(p)/(double)csvLength)*100);
//             statNtice.append(out_stringl);
//         }
//     }

//     ntice_pocet.clear();
//     ntice_typ_pocet.clear();

//     exportCsv(statNtice, pwd() + "StatistikaNtice_" + suborName() + ".csv");
// }

// void statArchiv(uint n,uint m, const Kombinacie &kombinacie){

//     CSV csv;
//     QStringList line;
//     QVector<Cislovacky> _cislovacky;

//     Kombinacia pred{};
//     for(auto &akt : kombinacie){
//         _cislovacky.append(cislovacky(akt, pred));
//         pred = akt;
//     }

//     //    Počet možných
//     line << "" << "Pocet moznych";
//     QVector<uint> pm(11,0);
//     for(uint cislo{1}; cislo <= m; ++cislo){
//         if(P(cislo)) ++pm[0];
//         if(N(cislo)) ++pm[1];
//         if(PR(cislo)) ++pm[2];
//         if(Mc(cislo)) ++pm[3];
//         if(Vc(cislo)) ++pm[4];
//         if(C19(cislo)) ++pm[5];
//         if(C0(cislo)) ++pm[6];
//         if(cC(cislo)) ++pm[7];
//         if(Cc(cislo)) ++pm[8];
//         if(CC(cislo)) ++pm[9];
//     }
//     pm[10] = n - 1;

//     for(auto el : pm){
//         line << uintToQString(el);
//     }
//     csv.append(line);

//     //    teoret možné
//     line.clear();
//     line << "" << "Teor mozne";
//     for(auto el : pm){
//         QString tm("0-");
//         if(el >= n)
//             tm += uintToQString(n);
//         else
//             tm += uintToQString(el);
//         line << tm;
//     }
//     csv.append(line);

//     QVector<QVector<double>> ap(11,{}), hp(11,{}), vp(11,{});

//     for(unsigned i{0}; i <= n; ++i){
//         csv.push_back({});
//         line.clear();

//         QVector<int> pu(11,0), krok(11,0), vyskyt(11,0), sums(11,0);
//         QVector<QVector<int>> krok_pocet(11,{});             // krok, pocet pre N,P,PR...

//         int riadok = 1;
//         for (auto &qv : _cislovacky) {
//             for(int poz{0}; poz < 11; ++poz){
//                 if(qv[poz] == i){
//                     ++pu[poz];                          // pocet udalost i
//                     vyskyt[poz] = riadok;
//                     if(krok[poz] > 0){
//                         krok_pocet[poz].push_back(krok[poz]);
//                         sums[poz] += krok[poz];
//                     }
//                     krok[poz] = 0;
//                 }
//                 else{
//                     krok[poz]++;
//                 }
//             }
//             ++riadok;
//         }

//         line << "" << "Pocet udalost " + uintToQString(i) ;
//         for(auto c : pu){
//             line << uintToQString(c);
//         }
//         csv.append(line);

//         //    Súčet diferencií krok udalosť i
//         line.clear();
//         line << "" << "Sucet diferencii krok udalost " + uintToQString(i);
//         for(int j{0}; j< 11; ++j){
//             line << uintToQString(sums[j]);
//         }
//         csv.append(line);

//         //    Krok aritmetický priemer krok udalosť  i
//         line.clear();
//         line << "" << "Krok aritmeticky priemer udalost " + uintToQString(i);
//         QVector<double> aritmeticky_priemer(11, 0.f);
//         for(int j{0}; j< 11; ++j){
//             if(krok_pocet[j].size() > 0)
//                 aritmeticky_priemer[j] = static_cast<double>(sums[j])/krok_pocet[j].size();
//             line << uintToQString(round(aritmeticky_priemer[j]));
//         }

//         csv.append(line);

//         //    Krok harmonický priemer krok udalosť i
//         line.clear();
//         line << "" << "Krok harmonicky priemer udalost " + uintToQString(i);
//         QVector<double> harmonicky_priemer(11, 0.f);
//         for(int j{0}; j< 11; ++j){
//             double hp = 0.f;

//             for(auto &c : krok_pocet[j]){
//                 hp += 1.0f/c;
//             }
//             if(hp > 0)
//                 harmonicky_priemer[j] = krok_pocet[j].size()/hp;
//             line << uintToQString(round(harmonicky_priemer[j]));
//         }

//         csv.append(line);

//         //    Krok vážený priemer krok udalosť i
//         QSet<int> set;
//         line.clear();
//         line << "" << "Krok vazeny priemer udalost " + uintToQString(i);
//         QVector<double> vazeny_priemer(11, 0.f);
//         for(int j{0}; j< 11; ++j){
//             double sum = 0.f;

//             for(auto el : krok_pocet[j])
//                 set.insert(el);
//             for(auto el : set)
//                 sum += el;

//             if(sum > 0)
//                 vazeny_priemer[j] = sums[j]/sum;
//             line << uintToQString(round(vazeny_priemer[j]));
//         }

//         csv.append(line);

//         //    Riadok posedný výskyt udalosť i
//         line.clear();
//         line << "" << "Riadok posledny vyskyt udalost " + uintToQString(i);
//         for(int j{0}; j< 11; ++j){
//             line << uintToQString(vyskyt[j]);
//         }
//         csv.append(line);

//         //    Riadokposedný výskyt i + Krok aritmetický priemer
//         line.clear();
//         line << "" << "Krok aritmeticky priemer + riadok posledny vyskyt " + uintToQString(i);
//         for(int j{0}; j< 11; ++j){
//             line << uintToQString(round(vyskyt[j] + aritmeticky_priemer[j]));
//             ap[j].push_back(round(vyskyt[j] + aritmeticky_priemer[j]));
//         }

//         csv.append(line);

//         //    Riadok posedný výskyt i + Krok harmonický priemer
//         line.clear();
//         line << "" << "Krok harmonicky priemer + riadok posledny vyskyt " + uintToQString(i);
//         for(int j{0}; j< 11; ++j){
//             line << uintToQString(round(vyskyt[j] + harmonicky_priemer[j]));
//             hp[j].push_back(round(vyskyt[j] + harmonicky_priemer[j]));
//         }

//         csv.append(line);

//         //    Riadok posedný výskyt i + Krok vážený priemer
//         line.clear();
//         line << "" << "Krok vazeny priemer + riadok posledny vyskyt " + uintToQString(i);
//         for(int j{0}; j< 11; ++j){
//             line << uintToQString(round(vyskyt[j] + vazeny_priemer[j]));
//             vp[j].push_back(vyskyt[j] + vazeny_priemer[j]);
//         }

//         csv.append(line);
//     }

//     auto najblizsia = [](QVector<double> &vec, unsigned dlzka){
//         double naj = dlzka + 1;
//         int i = 0, ret = -1;

//         for(auto c : vec){
//             auto val = abs(dlzka - c);
//             if(val < naj){
//                 naj = val;
//                 ret = i;
//             }
//             ++i;
//         }
//         return ret;
//     };

// //    csv.append({});
//     csv.push_back({});
//     line.clear();
//     line << "" << "Riadok posedny vyskyt + Krok aritmeticky priemer";
//     for(int i{0}; i < 11; ++i){
//         line << QString::number(najblizsia(ap[i], _cislovacky.size()));
//     }
//     csv.append(line);

//     line.clear();
//     line << "" << "Riadok posedny vyskyt + Krok harmonicky priemer";
//     for(int i{0}; i < 11; ++i){
//         line << QString::number(najblizsia(hp[i], _cislovacky.size()));
//     }
//     csv.append(line);

//     line.clear();
//     line << "" << "Riadok posedny vyskyt + Krok vazeny priemer";
//     for(int i{0}; i < 11; ++i){
//         line << QString::number(najblizsia(vp[i], _cislovacky.size()));
//     }
//     csv.append(line);

//     line.clear();
//     line << "" << "Riadok posedný výskyt - Kriterium v  r+1 (riadku)";
//     //    for(int i{0}; i < 11; ++i){
//     //        auto naj = najblizsia(najblizsia(hp[i], _cislovacky.size()));
//     //        if()
//     //    }
//     csv.append(line);

//     exportCsv(csv, pwd() + "StatistikaCislovacky_" + suborName() + ".csv");
// }
